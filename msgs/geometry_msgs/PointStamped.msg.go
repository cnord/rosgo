// Code generated by ros-gen-go.
// source: PointStamped.msg
// DO NOT EDIT!
package geometry_msgs

import (
	"io"

	"github.com/ppg/rosgo/ros"
	"github.com/ppg/rosgo/msgs/std_msgs"
)

type _MsgPointStamped struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgPointStamped) Text() string {
	return t.text
}

func (t *_MsgPointStamped) Name() string {
	return t.name
}

func (t *_MsgPointStamped) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgPointStamped) NewMessage() ros.Message {
	m := new(PointStamped)

	return m
}

var (
	MsgPointStamped = &_MsgPointStamped{
		`# This represents a Point with reference coordinate frame and timestamp
Header header
Point point
`,
		"geometry_msgs/PointStamped",
		"e948b3cf3f45aaeaedb063e8b966cf1f",
	}
)

type PointStamped struct {
	Header std_msgs.Header
	Point  Point
}

func (m *PointStamped) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "Point", &m.Point); err != nil {
		return err
	}

	return
}

func (m *PointStamped) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// Point
	if err = ros.DeserializeMessageField(r, "Point", &m.Point); err != nil {
		return err
	}

	return
}
