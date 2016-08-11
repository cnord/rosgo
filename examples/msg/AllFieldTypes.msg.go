// Code generated by ros-gen-go.
// source: AllFieldTypes.msg
// DO NOT EDIT!
package msgs

import (
	"io"
	"encoding/binary"
	"fmt"

	"github.com/ppg/rosgo/ros"
	"github.com/ppg/rosgo/msgs/std_msgs"
)

type _MsgAllFieldTypes struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgAllFieldTypes) Text() string {
	return t.text
}

func (t *_MsgAllFieldTypes) Name() string {
	return t.name
}

func (t *_MsgAllFieldTypes) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgAllFieldTypes) NewMessage() ros.Message {
	m := new(AllFieldTypes)

	return m
}

var (
	MsgAllFieldTypes = &_MsgAllFieldTypes{
		`byte FOO=1
byte BAR=2
string HOGE=hoge

Header h
byte b
int8 i8
int16 i16
int32 i32
int64 i64
uint8 u8
uint16 u16
uint32 u32
uint64 u64
float32 f32
float64 f64
time t
duration d
string s
std_msgs/ColorRGBA c
uint32[] dyn_ary
uint32[2] fix_ary
#std_msgs/ColorRGBA[] msg_ary
`,
		"msgs/AllFieldTypes",
		"d001ba01d11826b42fc4dbbd1ffacb2c",
	}
)

type AllFieldTypes struct {
	H      std_msgs.Header
	B      int8
	I8     int8
	I16    int16
	I32    int32
	I64    int64
	U8     uint8
	U16    uint16
	U32    uint32
	U64    uint64
	F32    float32
	F64    float64
	T      ros.Time
	D      ros.Duration
	S      string
	C      std_msgs.ColorRGBA
	DynAry []uint32
	FixAry [2]uint32
}

func (m *AllFieldTypes) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.H); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "byte", &m.B); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "int8", &m.I8); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "int16", &m.I16); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "int32", &m.I32); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "int64", &m.I64); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint8", &m.U8); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint16", &m.U16); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint32", &m.U32); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint64", &m.U64); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.F32); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float64", &m.F64); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "time", &m.T); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "duration", &m.D); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "string", &m.S); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "std_msgs/ColorRGBA", &m.C); err != nil {
		return err
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.DynAry)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.DynAry {
		if err = ros.SerializeMessageField(w, "uint32", &elem); err != nil {
			return err
		}
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.FixAry)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.FixAry {
		if err = ros.SerializeMessageField(w, "uint32", &elem); err != nil {
			return err
		}
	}

	return
}

func (m *AllFieldTypes) Deserialize(r io.Reader) (err error) {
	// H
	if err = ros.DeserializeMessageField(r, "Header", &m.H); err != nil {
		return err
	}

	// B
	if err = ros.DeserializeMessageField(r, "byte", &m.B); err != nil {
		return err
	}

	// I8
	if err = ros.DeserializeMessageField(r, "int8", &m.I8); err != nil {
		return err
	}

	// I16
	if err = ros.DeserializeMessageField(r, "int16", &m.I16); err != nil {
		return err
	}

	// I32
	if err = ros.DeserializeMessageField(r, "int32", &m.I32); err != nil {
		return err
	}

	// I64
	if err = ros.DeserializeMessageField(r, "int64", &m.I64); err != nil {
		return err
	}

	// U8
	if err = ros.DeserializeMessageField(r, "uint8", &m.U8); err != nil {
		return err
	}

	// U16
	if err = ros.DeserializeMessageField(r, "uint16", &m.U16); err != nil {
		return err
	}

	// U32
	if err = ros.DeserializeMessageField(r, "uint32", &m.U32); err != nil {
		return err
	}

	// U64
	if err = ros.DeserializeMessageField(r, "uint64", &m.U64); err != nil {
		return err
	}

	// F32
	if err = ros.DeserializeMessageField(r, "float32", &m.F32); err != nil {
		return err
	}

	// F64
	if err = ros.DeserializeMessageField(r, "float64", &m.F64); err != nil {
		return err
	}

	// T
	if err = ros.DeserializeMessageField(r, "time", &m.T); err != nil {
		return err
	}

	// D
	if err = ros.DeserializeMessageField(r, "duration", &m.D); err != nil {
		return err
	}

	// S
	if err = ros.DeserializeMessageField(r, "string", &m.S); err != nil {
		return err
	}

	// C
	if err = ros.DeserializeMessageField(r, "std_msgs/ColorRGBA", &m.C); err != nil {
		return err
	}

	// DynAry
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for DynAry: %s", err)
		}
		m.DynAry = make([]uint32, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "uint32", &m.DynAry[i]); err != nil {
				return err
			}
		}
	}

	// FixAry
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for FixAry: %s", err)
		}
		if size > 2 {
			return fmt.Errorf("array size for FixAry too large: expected=2, got=%d", size)
		}
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "uint32", &m.FixAry[i]); err != nil {
				return err
			}
		}
	}

	return
}
