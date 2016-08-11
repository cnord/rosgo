// Code generated by ros-gen-go.
// source: Int8.msg
// DO NOT EDIT!
package std_msgs

import (
	"io"

	"github.com/ppg/rosgo/ros"
)

type _MsgInt8 struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgInt8) Text() string {
	return t.text
}

func (t *_MsgInt8) Name() string {
	return t.name
}

func (t *_MsgInt8) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgInt8) NewMessage() ros.Message {
	m := new(Int8)

	return m
}

var (
	MsgInt8 = &_MsgInt8{
		`int8 data
`,
		"std_msgs/Int8",
		"09cdd02665af5f0d32bec3c40d0e8f62",
	}
)

type Int8 struct {
	Data int8
}

func (m *Int8) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "int8", &m.Data); err != nil {
		return err
	}

	return
}

func (m *Int8) Deserialize(r io.Reader) (err error) {
	// Data
	if err = ros.DeserializeMessageField(r, "int8", &m.Data); err != nil {
		return err
	}

	return
}
