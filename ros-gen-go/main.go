package main

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"fmt"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"unicode"

	flag "github.com/ogier/pflag"
)

var (
	uribase     string = "github.com/cnord"
	infile      string
	outfile     string
	packageName string
	dryRun      bool
)

type loopVarSetter interface {
	SetLoopVar(interface{})
}

func readFile(infile string) ([]byte, error) {
	// That gimmick used to close infile before outfile,
	// it takes off a race condition when outfile could be closed before infile, so,
	// it means that make util could regenerate .msg.go files
	finfile, err := os.Open(infile)
	if err != nil {
		//log.Fatalf("failed to open infile %s: %s", infile, err)
		return nil, err;
	}
	defer finfile.Close()

	var finstat os.FileInfo
	finstat, err = finfile.Stat()
	if err != nil {
		//log.Fatalf("failed during receiving infile stat %s: %s", infile, err)
		return nil, err;
	}

	var fisz int;
	indata := make([]byte, finstat.Size())
	fisz, err = finfile.Read(indata)
	if (err != nil) || (finstat.Size() != int64(fisz)) {
		//log.Fatalf("failed to read infile %s: %s", infile, err)
		return nil, err;
	}
	return indata, nil
}

func main() {
	log.SetFlags(0)

	flag.StringVarP(&infile, "in", "i", "", "input file")
	flag.StringVarP(&outfile, "out", "o", "", "output file; defaults to '<input file>.go'")
	flag.StringVarP(&packageName, "package", "p", "", "package name for generated file; defaults to 'msgs' or 'srvs'")
	flag.BoolVar(&dryRun, "dry-run", false, "output the file that would be generated to stdout")
	flag.Parse()

	if flag.NArg() < 1 {
		log.Printf("must provide type of generator")
		flag.PrintDefaults()
		os.Exit(1)
	}
	templateType := flag.Arg(0)
	if packageName == "" {
		packageName = templateType + "s"
	}

	basename := fmt.Sprintf("%s.tmpl", templateType)
	data, err := Asset(basename)
	if err != nil {
		log.Printf("unrecognized generator template: %s (%s)", templateType, err)
		flag.PrintDefaults()
		os.Exit(1)
	}
	tmpl := template.New(basename)
	tmpl = tmpl.Funcs(map[string]interface{}{
		// HACK(ppg): Allow setting a loop variable a struct so we can use it
		"setloopvar": func(setter loopVarSetter, value interface{}) interface{} {
			setter.SetLoopVar(value)
			return setter
		},
	})
	tmpl, err = tmpl.Parse(string(data))
	if err != nil {
		log.Printf("unable to template %s: %s", templateType, err)
		os.Exit(1)
	}

	data, err = Asset("msg.partial.tmpl")
	if err != nil {
		log.Printf("unrecognized generator template: %s (%s)", templateType, err)
		flag.PrintDefaults()
		os.Exit(1)
	}
	tmpl2 := tmpl.New("msg.partial.tmpl")
	_, err = tmpl2.Parse(string(data))
	if err != nil {
		log.Printf("unable to template %s: %s", templateType, err)
		os.Exit(1)
	}

	if flag.NArg() > 1 {
		log.Printf("unrecognized arguments: %v", flag.Args()[1:])
		flag.PrintDefaults()
		os.Exit(1)
	}

	if infile == "" {
		log.Printf("must provide input file")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if outfile == "" {
		outfile = infile + ".go"
	}

	// Read input file
	data, err = readFile(infile)
	if err != nil {
		log.Fatalf("failed to read infile %s: %s", infile, err)
	}
	basename = filepath.Base(infile)
	fileInfo := FileInfo{
		InFile:      infile,
		InFileBase:  filepath.Base(infile),
		Raw:         string(data),
		MD5Sum:      fmt.Sprintf("%x", md5.Sum(data)),
		PackageName: packageName,
		Name:        strings.TrimSuffix(basename, filepath.Ext(basename)),
	}

	// Parse by type
	var spec interface{}
	switch templateType {
	case "msg":
		var msgSpec *MsgSpec
		msgSpec, err = parseMsgSpec(fileInfo.PackageName, fileInfo.Name, data)
		if err != nil {
			log.Fatalf("failed to parse %s spec: %s", templateType, err)
		}
		spec = msgSpec

	case "srv":
		var srvSpec *SrvSpec
		srvSpec, err = parseSrvSpec(fileInfo.PackageName, fileInfo.Name, data)
		if err != nil {
			log.Fatalf("failed to parse %s spec: %s", templateType, err)
		}
		spec = srvSpec

	default:
		log.Fatalf("no parser configured: %s", templateType)
	}

	buf := bytes.NewBuffer([]byte{})
	err = tmpl.Execute(buf, map[string]interface{}{"FileInfo": fileInfo, "Spec": spec})
	if err != nil {
		log.Fatalf("failed to generate Go file: %s", err)
	}

	fset := token.NewFileSet()
	ast, err := parser.ParseFile(fset, "", buf.Bytes(), parser.ParseComments)
	if err != nil {
		log.Fatalf("bad Go source code was generated: %s\n%s", err, buf.String())
	}
	buf.Reset()
	err = (&printer.Config{Mode: printer.TabIndent | printer.UseSpaces, Tabwidth: 8}).Fprint(buf, fset, ast)
	if err != nil {
		log.Fatalf("generated Go source code could not be reformatted: %s", err)
	}

	if dryRun {
		fmt.Println(buf.String())
		return
	}

	err = ioutil.WriteFile(outfile, buf.Bytes(), 0644)
	if err != nil {
		log.Fatalf("failed to write go file: %s", err)
	}
	log.Printf("Wrote %s from %s", outfile, infile)
}

type FileInfo struct {
	InFile      string
	InFileBase  string
	Raw         string
	MD5Sum      string
	PackageName string
	Name        string
}

type MsgSpec struct {
	Raw         string
	MD5Sum      string
	PackageName string
	Name        string

	Fields     []*msgField
	Constants  []*msgConstant
	HasBuiltIn bool
	HasSlice   bool
	HasArray   bool

	packageMap map[string]struct{}
}

func (m MsgSpec) Packages() (ret []string) {
	ret = make([]string, 0, len(m.packageMap))
	for k := range m.packageMap {
		ret = append(ret, k)
	}
	sort.StringSlice(ret).Sort()
	return
}

//byte FOO=1
//byte BAR=2
//string HOGE=hoge
var constMatcher = regexp.MustCompile(`^([\w/]+)\s+(\w+)\s*=\s*(\w+)`)

//Header h # go:package=github.com/cnord/rosgo/msgs/std_msgs
//byte b
//std_msgs/ColorRGBA c
//uint32[] dyn_ary
//uint32[2] fix_ary
//std_msgs/ColorRGBA[] msg_ary
var fieldMatcher = regexp.MustCompile(`^([\w/]+)(\[(\d*)\])?\s+(\w+)`)

// go:package=github.com/cnord/rosgo/msgs/std_msgs
var goOptionMatcher = regexp.MustCompile(`go:(\w+)=([^\s]+)`)

//uint8 status
//uint8 PENDING         = 0   # The goal has yet to be processed by the action server
func parseMsgSpec(packageName, name string, data []byte) (*MsgSpec, error) {
	spec := new(MsgSpec)
	spec.Raw = string(data)
	md5sum, err := exec.Command("rosmsg", "md5", fmt.Sprintf("%s/%s", packageName, name)).Output()
	if err != nil {
		return nil, err
	}
	spec.MD5Sum = strings.Trim(string(md5sum), " \n")
	spec.PackageName = packageName
	spec.Name = name

	spec.packageMap = map[string]struct{}{uribase+"/rosgo/ros": struct{}{}}

	// Read the data line by line
	buf := bytes.NewBuffer(data)
	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		// Check constant first
		items := constMatcher.FindStringSubmatch(line)
		if items != nil {
			//log.Println("constant", line)
			//log.Printf("items(%d): %+v", len(items), items)
			//for i := 0; i < len(items); i++ {
			//	log.Printf("items[%d]: %+v", i, items[i])
			//}
			//items(4): [string HOGE=hoge string HOGE hoge]
			//items[0]: string HOGE=hoge
			//items[1]: string
			//items[2]: HOGE
			//items[3]: hoge
			rosName := items[2]
			rosType := items[1]
			rosValue := items[3]
			constant := newMsgConstant(rosName, rosType, rosValue)
			//log.Printf("constant: %+v", constant)
			spec.Constants = append(spec.Constants, constant)
			continue
		}

		// Now check field
		items = fieldMatcher.FindStringSubmatch(line)
		if items != nil {
			//log.Println("field", line)
			//log.Printf("items(%d): %+v", len(items), items)
			//for i := 0; i < len(items); i++ {
			//	log.Printf("items[%d]: %+v", i, items[i])
			//}
			// Parse go options
			goOptions := make(map[string]string)
			for _, option := range goOptionMatcher.FindAllStringSubmatch(line, -1) {
				goOptions[option[1]] = option[2]
			}
			//log.Println("goOptions", goOptions)
			//items(5): [uint32[2] fix_ary uint32 [2] 2 fix_ary]
			//items[0]: uint32[2] fix_ary
			//items[1]: uint32
			//items[2]: [2]
			//items[3]: 2
			//items[4]: fix_ary
			rosName := items[4]
			rosType := items[1]
			isSliceOrArray := len(items[2]) > 0
			arraySize := items[3]
			field := newMsgField(rosName, rosType, isSliceOrArray, arraySize)
			if field.BuiltIn {
				spec.HasBuiltIn = true
			}
			if field.IsArray {
				spec.HasSlice = true
				if field.ArraySize > 0 {
					spec.HasArray = true
				}
			}
			if field.GoImportName != "" && field.GoImportName != packageName {
				if pkg, ok := goOptions["package"]; ok {
					spec.packageMap[pkg] = struct{}{}
				} else if pkg, ok := builtInImports[field.GoImportName]; ok {
					spec.packageMap[pkg] = struct{}{}
				} else {
					log.Printf("WARNING: go import name not found in go:package option nor built-in imports: %s", field.GoImportName)
				}
			}
			//log.Printf("field: %+v", field)
			spec.Fields = append(spec.Fields, field)
			continue
		}

		if len(line) != 0 && !strings.HasPrefix(line, "#") {
			return nil, fmt.Errorf("unrecognized msg line: %s", line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return spec, nil
}

type msgField struct {
	IsArray      bool
	ArraySize    int
	Name         string
	TypeName     string
	BuiltIn      bool
	GoImportName string
	GoTypeName   string
	GoZeroValue  string

	LoopVar interface{}
}

func (m *msgField) SetLoopVar(i interface{}) {
	m.LoopVar = i
}

func newMsgField(rosName, rosType string, isArray bool, arraySize string) (field *msgField) {
	//log.Printf("rosName: %s", rosName)
	//log.Printf("rosType: %s", rosType)
	field = new(msgField)
	field.Name = snakeToCamel(rosName)
	field.TypeName = rosType
	field.IsArray = isArray
	if isArray && len(arraySize) > 0 {
		var err error
		field.ArraySize, err = strconv.Atoi(arraySize)
		if err != nil {
			panic(err)
		}
	}

	// Try to get field info from built-in first
	info, ok := builtInInfo[rosType]
	if ok {
		field.BuiltIn = true
		field.GoImportName = ""
		field.GoTypeName = info.TypeName
		field.GoZeroValue = info.ZeroValue
		return
	}

	// Not built-in, deconstruct
	field.BuiltIn = false

	// TODO(ppg): Can we set all complex types to nil instead?

	// Special case Header per ROS docs
	if rosType == "Header" {
		field.GoImportName = "std_msgs"
		field.GoTypeName = "std_msgs.Header"
		field.GoZeroValue = "std_msgs.Header{}"
		return
	}

	// Split on / for namespaced messages
	items := strings.Split(rosType, "/")
	goMessageTypeName := snakeToCamel(items[0])

	// If we have a namespace adjust goMessageName
	if len(items) == 2 {
		field.GoImportName = items[0]
		goMessageTypeName = snakeToCamel(items[1])
		// if its not our package then prefix and return
		if field.GoImportName != packageName {
			field.GoTypeName = fmt.Sprintf("%s.%s", field.GoImportName, goMessageTypeName)
			field.GoZeroValue = fmt.Sprintf("%s.%s{}", field.GoImportName, goMessageTypeName)
			return
		}
	}

	// Returned non-namespaced
	field.GoTypeName = fmt.Sprintf("%s", goMessageTypeName)
	field.GoZeroValue = fmt.Sprintf("%s{}", goMessageTypeName)
	return
}

type msgConstant struct {
	Name     string
	TypeName string
	Value    string
}

func newMsgConstant(rosName, rosType, value string) (constant *msgConstant) {
	constant = new(msgConstant)
	constant.Name = snakeToCamel(rosName)
	constant.TypeName = rosType
	constant.Value = value
	return
}

type SrvSpec struct {
	Raw         string
	MD5Sum      string
	PackageName string
	Name        string

	RequestSpec  *MsgSpec
	ResponseSpec *MsgSpec

	packageMap map[string]struct{}
}

func (s SrvSpec) Packages() (ret []string) {
	ret = make([]string, 0, len(s.packageMap))
	for k := range s.packageMap {
		ret = append(ret, k)
	}
	sort.StringSlice(ret).Sort()
	return
}

//int32 a
//int32 b
//---
//int32 sum
func parseSrvSpec(packageName, name string, data []byte) (*SrvSpec, error) {
	spec := new(SrvSpec)
	spec.Raw = string(data)
	md5sum, err := exec.Command("rossrv", "md5", fmt.Sprintf("%s/%s", packageName, name)).Output()
	if err != nil {
		return nil, err
	}
	spec.MD5Sum = strings.Trim(string(md5sum), " \n")
	spec.PackageName = packageName
	spec.Name = name

	spec.packageMap = map[string]struct{}{uribase+"/rosgo/ros": struct{}{}}

	raws := strings.Split(string(data), "---")
	spec.RequestSpec, err = parseMsgSpec(packageName, fmt.Sprintf("%sRequest", name), []byte(raws[0]))
	if err != nil {
		return nil, err
	}
	spec.ResponseSpec, err = parseMsgSpec(packageName, fmt.Sprintf("%sResponse", name), []byte(raws[1]))
	if err != nil {
		return nil, err
	}

	// Collapse request/response packages into this map
	for k := range spec.RequestSpec.packageMap {
		spec.packageMap[k] = struct{}{}
	}
	for k := range spec.ResponseSpec.packageMap {
		spec.packageMap[k] = struct{}{}
	}

	return spec, nil
}

type goInfo struct {
	TypeName  string
	ZeroValue string
}

var builtInInfo = map[string]goInfo{
	"bool": goInfo{"bool", "false"},
	"byte": goInfo{"int8", "0"},
	"char": goInfo{"uint8", "0"},
	// TODO(ppg): Change to time.Duration, handle conversion to ROS type during serialization/deserialization
	"duration": goInfo{"ros.Duration", "ros.Duration{}"},
	"float32":  goInfo{"float32", "0.0"},
	"float64":  goInfo{"float64", "0.0"},
	"int8":     goInfo{"int8", "0"},
	"int16":    goInfo{"int16", "0"},
	"int32":    goInfo{"int32", "0"},
	"int64":    goInfo{"int64", "0"},
	"string":   goInfo{"string", `""`},
	// TODO(ppg): Change to time.Time, handle conversion to ROS type during serialization/deserialization
	"time":   goInfo{"ros.Time", "ros.Time{}"},
	"uint8":  goInfo{"uint8", "0"},
	"uint16": goInfo{"uint16", "0"},
	"uint32": goInfo{"uint32", "0"},
	"uint64": goInfo{"uint64", "0"},
}

var builtInImports = map[string]string{
	"actionlib_msgs":     uribase+"/rosgo/msgs/actionlib_msgs",
	"common_msgs":        uribase+"/rosgo/msgs/common_msgs",
	"control_msgs":       uribase+"/rosgo/msgs/control_msgs",
	"diagnostic_msgs":    uribase+"/rosgo/msgs/diagnostic_msgs",
	"geographic_msgs":    uribase+"/rosgo/msgs/geographic_msgs",
	"geometry_msgs":      uribase+"/rosgo/msgs/geometry_msgs",
	"map_msgs":           uribase+"/rosgo/msgs/map_msgs",
	"mavros_msgs":        uribase+"/rosgo/msgs/mavros_msgs",
	"nav_msgs":           uribase+"/rosgo/msgs/nav_msgs",
	"rosgraph_msgs":      uribase+"/rosgo/msgs/rosgraph_msgs",
	"sensor_msgs":        uribase+"/rosgo/msgs/sensor_msgs",
	"shape_msgs":         uribase+"/rosgo/msgs/shape_msgs",
	"smach_msgs":         uribase+"/rosgo/msgs/smach_msgs",
	"std_msgs":           uribase+"/rosgo/msgs/std_msgs",
	"stereo_msgs":        uribase+"/rosgo/msgs/stereo_msgs",
	"tf2_msgs":           uribase+"/rosgo/msgs/tf2_msgs",
	"trajectory_msgs":    uribase+"/rosgo/msgs/trajectory_msgs",
	"visualization_msgs": uribase+"/rosgo/msgs/visualization_msgs",
	"uuid_msgs":          uribase+"/rosgo/msgs/uuid_msgs",
}

// snakeToCamel returns a string converted from snake case to uppercase
func snakeToCamel(s string) string {
	var result string

	words := strings.Split(s, "_")

	for _, word := range words {
		if upper := strings.ToUpper(word); commonInitialisms[upper] {
			result += upper
			continue
		}

		if len(word) > 0 {
			w := []rune(word)
			w[0] = unicode.ToUpper(w[0])
			result += string(w)
		}
	}

	return result
}

// commonInitialisms, taken from
// https://github.com/golang/lint/blob/32a87160691b3c96046c0c678fe57c5bef761456/lint.go#L702
var commonInitialisms = map[string]bool{
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SQL":   true,
	"SSH":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XSRF":  true,
	"XSS":   true,
}
