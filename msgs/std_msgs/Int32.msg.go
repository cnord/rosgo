// Code generated by ros-gen-go.
// source: Int32.msg
// DO NOT EDIT!
package std_msgs

import (
	"io"

	"github.com/cnord/rosgo/ros"
)

type _MsgInt32 struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgInt32) Text() string {
	return t.text
}

func (t *_MsgInt32) Name() string {
	return t.name
}

func (t *_MsgInt32) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgInt32) NewMessage() ros.Message {
	m := new(Int32)

	return m
}

var (
	MsgInt32 = &_MsgInt32{
		`int32 data`,
		"std_msgs/Int32",
		"da5909fbe378aeaf85e547e830cc1bb7",
	}
)

type Int32 struct {
	Data int32
}

func (m *Int32) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "int32", &m.Data); err != nil {
		return err
	}

	return
}

func (m *Int32) Deserialize(r io.Reader) (err error) {
	// Data
	if err = ros.DeserializeMessageField(r, "int32", &m.Data); err != nil {
		return err
	}

	return
}
