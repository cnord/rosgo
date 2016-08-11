// Code generated by ros-gen-go.
// source: PolygonStamped.msg
// DO NOT EDIT!
package geometry_msgs

import (
	"io"

	"github.com/ppg/rosgo/msgs/std_msgs"
	"github.com/ppg/rosgo/ros"
)

type _MsgPolygonStamped struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgPolygonStamped) Text() string {
	return t.text
}

func (t *_MsgPolygonStamped) Name() string {
	return t.name
}

func (t *_MsgPolygonStamped) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgPolygonStamped) NewMessage() ros.Message {
	m := new(PolygonStamped)

	return m
}

var (
	MsgPolygonStamped = &_MsgPolygonStamped{
		`# This represents a Polygon with reference coordinate frame and timestamp
Header header
Polygon polygon
`,
		"geometry_msgs/PolygonStamped",
		"aa0712e581b3057dba582c695ff17d61",
	}
)

type PolygonStamped struct {
	Header  std_msgs.Header
	Polygon Polygon
}

func (m *PolygonStamped) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "Polygon", &m.Polygon); err != nil {
		return err
	}

	return
}

func (m *PolygonStamped) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// Polygon
	if err = ros.DeserializeMessageField(r, "Polygon", &m.Polygon); err != nil {
		return err
	}

	return
}