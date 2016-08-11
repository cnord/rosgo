// Code generated by ros-gen-go.
// source: Transform.msg
// DO NOT EDIT!
package geometry_msgs

import (
	"io"

	"github.com/ppg/rosgo/ros"
)

type _MsgTransform struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgTransform) Text() string {
	return t.text
}

func (t *_MsgTransform) Name() string {
	return t.name
}

func (t *_MsgTransform) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgTransform) NewMessage() ros.Message {
	m := new(Transform)

	return m
}

var (
	MsgTransform = &_MsgTransform{
		`# This represents the transform between two coordinate frames in free space.

Vector3 translation
Quaternion rotation
`,
		"geometry_msgs/Transform",
		"756be060b1c8cf0e64a10ba16909d887",
	}
)

type Transform struct {
	Translation Vector3
	Rotation    Quaternion
}

func (m *Transform) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Vector3", &m.Translation); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "Quaternion", &m.Rotation); err != nil {
		return err
	}

	return
}

func (m *Transform) Deserialize(r io.Reader) (err error) {
	// Translation
	if err = ros.DeserializeMessageField(r, "Vector3", &m.Translation); err != nil {
		return err
	}

	// Rotation
	if err = ros.DeserializeMessageField(r, "Quaternion", &m.Rotation); err != nil {
		return err
	}

	return
}
