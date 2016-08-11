// Code generated by ros-gen-go.
// source: Log.msg
// DO NOT EDIT!
package rosgraph_msgs

import (
	"io"
	"encoding/binary"
	"fmt"

	"github.com/ppg/rosgo/ros"
	"github.com/ppg/rosgo/msgs/std_msgs"
)

type _MsgLog struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgLog) Text() string {
	return t.text
}

func (t *_MsgLog) Name() string {
	return t.name
}

func (t *_MsgLog) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgLog) NewMessage() ros.Message {
	m := new(Log)

	return m
}

var (
	MsgLog = &_MsgLog{
		`##
## Severity level constants
##
byte DEBUG=1 #debug level
byte INFO=2  #general level
byte WARN=4  #warning level
byte ERROR=8 #error level
byte FATAL=16 #fatal/critical level
##
## Fields
##
Header header
byte level
string name # name of the node
string msg # message 
string file # file the message came from
string function # function the message came from
uint32 line # line the message came from
string[] topics # topic names that the node publishes
`,
		"rosgraph_msgs/Log",
		"7e4c0d13ed75610b7a2bb18594155bf8",
	}
)

type Log struct {
	Header   std_msgs.Header
	Level    int8
	Name     string
	Msg      string
	File     string
	Function string
	Line     uint32
	Topics   []string
}

func (m *Log) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "byte", &m.Level); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "string", &m.Name); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "string", &m.Msg); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "string", &m.File); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "string", &m.Function); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint32", &m.Line); err != nil {
		return err
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Topics)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Topics {
		if err = ros.SerializeMessageField(w, "string", &elem); err != nil {
			return err
		}
	}

	return
}

func (m *Log) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// Level
	if err = ros.DeserializeMessageField(r, "byte", &m.Level); err != nil {
		return err
	}

	// Name
	if err = ros.DeserializeMessageField(r, "string", &m.Name); err != nil {
		return err
	}

	// Msg
	if err = ros.DeserializeMessageField(r, "string", &m.Msg); err != nil {
		return err
	}

	// File
	if err = ros.DeserializeMessageField(r, "string", &m.File); err != nil {
		return err
	}

	// Function
	if err = ros.DeserializeMessageField(r, "string", &m.Function); err != nil {
		return err
	}

	// Line
	if err = ros.DeserializeMessageField(r, "uint32", &m.Line); err != nil {
		return err
	}

	// Topics
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Topics: %s", err)
		}
		m.Topics = make([]string, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "string", &m.Topics[i]); err != nil {
				return err
			}
		}
	}

	return
}
