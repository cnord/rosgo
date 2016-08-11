// Code generated by ros-gen-go.
// source: Point.msg
// DO NOT EDIT!
package geometry_msgs

import (
	"io"

	"github.com/ppg/rosgo/ros"
)

type _MsgPoint struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgPoint) Text() string {
	return t.text
}

func (t *_MsgPoint) Name() string {
	return t.name
}

func (t *_MsgPoint) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgPoint) NewMessage() ros.Message {
	m := new(Point)

	return m
}

var (
	MsgPoint = &_MsgPoint{
		`# This contains the position of a point in free space
float64 x
float64 y
float64 z
`,
		"geometry_msgs/Point",
		"243a4e5999bbb6bb311c053eea9e3ead",
	}
)

type Point struct {
	X float64
	Y float64
	Z float64
}

func (m *Point) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "float64", &m.X); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float64", &m.Y); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float64", &m.Z); err != nil {
		return err
	}

	return
}

func (m *Point) Deserialize(r io.Reader) (err error) {
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

	return
}