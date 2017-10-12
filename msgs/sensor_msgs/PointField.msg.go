// Code generated by ros-gen-go.
// source: PointField.msg
// DO NOT EDIT!
package sensor_msgs

import (
	"io"

	"github.com/cnord/rosgo/ros"
)

type _MsgPointField struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgPointField) Text() string {
	return t.text
}

func (t *_MsgPointField) Name() string {
	return t.name
}

func (t *_MsgPointField) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgPointField) NewMessage() ros.Message {
	m := new(PointField)

	return m
}

var (
	MsgPointField = &_MsgPointField{
		`# This message holds the description of one point entry in the
# PointCloud2 message format.
uint8 INT8    = 1
uint8 UINT8   = 2
uint8 INT16   = 3
uint8 UINT16  = 4
uint8 INT32   = 5
uint8 UINT32  = 6
uint8 FLOAT32 = 7
uint8 FLOAT64 = 8

string name      # Name of field
uint32 offset    # Offset from start of point struct
uint8  datatype  # Datatype enumeration, see above
uint32 count     # How many elements in the field
`,
		"sensor_msgs/PointField",
		"268eacb2962780ceac86cbd17e328150",
	}
)

type PointField struct {
	Name     string
	Offset   uint32
	Datatype uint8
	Count    uint32
}

func (m *PointField) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "string", &m.Name); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint32", &m.Offset); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint8", &m.Datatype); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint32", &m.Count); err != nil {
		return err
	}

	return
}

func (m *PointField) Deserialize(r io.Reader) (err error) {
	// Name
	if err = ros.DeserializeMessageField(r, "string", &m.Name); err != nil {
		return err
	}

	// Offset
	if err = ros.DeserializeMessageField(r, "uint32", &m.Offset); err != nil {
		return err
	}

	// Datatype
	if err = ros.DeserializeMessageField(r, "uint8", &m.Datatype); err != nil {
		return err
	}

	// Count
	if err = ros.DeserializeMessageField(r, "uint32", &m.Count); err != nil {
		return err
	}

	return
}
