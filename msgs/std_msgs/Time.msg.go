// Code generated by ros-gen-go.
// source: Time.msg
// DO NOT EDIT!
package std_msgs

import (
	"io"

	"github.com/cnord/rosgo/ros"
)

type _MsgTime struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgTime) Text() string {
	return t.text
}

func (t *_MsgTime) Name() string {
	return t.name
}

func (t *_MsgTime) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgTime) NewMessage() ros.Message {
	m := new(Time)

	return m
}

var (
	MsgTime = &_MsgTime{
		`time data
`,
		"std_msgs/Time",
		"cd7166c74c552c311fbcc2fe5a7bc289",
	}
)

type Time struct {
	Data ros.Time
}

func (m *Time) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "time", &m.Data); err != nil {
		return err
	}

	return
}

func (m *Time) Deserialize(r io.Reader) (err error) {
	// Data
	if err = ros.DeserializeMessageField(r, "time", &m.Data); err != nil {
		return err
	}

	return
}
