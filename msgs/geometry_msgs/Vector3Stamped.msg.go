// Code generated by ros-gen-go.
// source: Vector3Stamped.msg
// DO NOT EDIT!
package geometry_msgs

import (
	"io"

	"github.com/cnord/rosgo/msgs/std_msgs"
	"github.com/cnord/rosgo/ros"
)

type _MsgVector3Stamped struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgVector3Stamped) Text() string {
	return t.text
}

func (t *_MsgVector3Stamped) Name() string {
	return t.name
}

func (t *_MsgVector3Stamped) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgVector3Stamped) NewMessage() ros.Message {
	m := new(Vector3Stamped)

	return m
}

var (
	MsgVector3Stamped = &_MsgVector3Stamped{
		`# This represents a Vector3 with reference coordinate frame and timestamp
Header header
Vector3 vector
`,
		"geometry_msgs/Vector3Stamped",
		"7b324c7325e683bf02a9b14b01090ec7",
	}
)

type Vector3Stamped struct {
	Header std_msgs.Header
	Vector Vector3
}

func (m *Vector3Stamped) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "Vector3", &m.Vector); err != nil {
		return err
	}

	return
}

func (m *Vector3Stamped) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// Vector
	if err = ros.DeserializeMessageField(r, "Vector3", &m.Vector); err != nil {
		return err
	}

	return
}
