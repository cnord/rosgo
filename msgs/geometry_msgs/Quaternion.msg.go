// Code generated by ros-gen-go.
// source: Quaternion.msg
// DO NOT EDIT!
package geometry_msgs

import (
	"io"

	"github.com/ppg/rosgo/ros"
)

type _MsgQuaternion struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgQuaternion) Text() string {
	return t.text
}

func (t *_MsgQuaternion) Name() string {
	return t.name
}

func (t *_MsgQuaternion) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgQuaternion) NewMessage() ros.Message {
	m := new(Quaternion)

	return m
}

var (
	MsgQuaternion = &_MsgQuaternion{
		`# This represents an orientation in free space in quaternion form.

float64 x
float64 y
float64 z
float64 w
`,
		"geometry_msgs/Quaternion",
		"4c29a42d17cc0d0a6ead2e4413ed0e21",
	}
)

type Quaternion struct {
	X float64
	Y float64
	Z float64
	W float64
}

func (m *Quaternion) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "float64", &m.X); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float64", &m.Y); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float64", &m.Z); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float64", &m.W); err != nil {
		return err
	}

	return
}

func (m *Quaternion) Deserialize(r io.Reader) (err error) {
	// X
	if err = ros.DeserializeMessageField(r, "float64", &m.X); err != nil {
		return err
	}

	// Y
	if err = ros.DeserializeMessageField(r, "float64", &m.Y); err != nil {
		return err
	}

	// Z
	if err = ros.DeserializeMessageField(r, "float64", &m.Z); err != nil {
		return err
	}

	// W
	if err = ros.DeserializeMessageField(r, "float64", &m.W); err != nil {
		return err
	}

	return
}
