// Code generated by ros-gen-go.
// source: PointCloud2.msg
// DO NOT EDIT!
package sensor_msgs

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/cnord/rosgo/msgs/std_msgs"
	"github.com/cnord/rosgo/ros"
)

type _MsgPointCloud2 struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgPointCloud2) Text() string {
	return t.text
}

func (t *_MsgPointCloud2) Name() string {
	return t.name
}

func (t *_MsgPointCloud2) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgPointCloud2) NewMessage() ros.Message {
	m := new(PointCloud2)

	return m
}

var (
	MsgPointCloud2 = &_MsgPointCloud2{
		`# This message holds a collection of N-dimensional points, which may
# contain additional information such as normals, intensity, etc. The
# point data is stored as a binary blob, its layout described by the
# contents of the "fields" array.

# The point cloud data may be organized 2d (image-like) or 1d
# (unordered). Point clouds organized as 2d images may be produced by
# camera depth sensors such as stereo or time-of-flight.

# Time of sensor data acquisition, and the coordinate frame ID (for 3d
# points).
Header header

# 2D structure of the point cloud. If the cloud is unordered, height is
# 1 and width is the length of the point cloud.
uint32 height
uint32 width

# Describes the channels and their layout in the binary data blob.
PointField[] fields

bool    is_bigendian # Is this data bigendian?
uint32  point_step   # Length of a point in bytes
uint32  row_step     # Length of a row in bytes
uint8[] data         # Actual point data, size is (row_step*height)

bool is_dense        # True if there are no invalid points
`,
		"sensor_msgs/PointCloud2",
		"1158d486dd51d683ce2f1be655c3c181",
	}
)

type PointCloud2 struct {
	Header      std_msgs.Header
	Height      uint32
	Width       uint32
	Fields      []PointField
	IsBigendian bool
	PointStep   uint32
	RowStep     uint32
	Data        []uint8
	IsDense     bool
}

func (m *PointCloud2) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint32", &m.Height); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint32", &m.Width); err != nil {
		return err
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Fields)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Fields {
		if err = ros.SerializeMessageField(w, "PointField", &elem); err != nil {
			return err
		}
	}

	if err = ros.SerializeMessageField(w, "bool", &m.IsBigendian); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint32", &m.PointStep); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint32", &m.RowStep); err != nil {
		return err
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Data)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Data {
		if err = ros.SerializeMessageField(w, "uint8", &elem); err != nil {
			return err
		}
	}

	if err = ros.SerializeMessageField(w, "bool", &m.IsDense); err != nil {
		return err
	}

	return
}

func (m *PointCloud2) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// Height
	if err = ros.DeserializeMessageField(r, "uint32", &m.Height); err != nil {
		return err
	}

	// Width
	if err = ros.DeserializeMessageField(r, "uint32", &m.Width); err != nil {
		return err
	}

	// Fields
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Fields: %s", err)
		}
		m.Fields = make([]PointField, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "PointField", &m.Fields[i]); err != nil {
				return err
			}
		}
	}

	// IsBigendian
	if err = ros.DeserializeMessageField(r, "bool", &m.IsBigendian); err != nil {
		return err
	}

	// PointStep
	if err = ros.DeserializeMessageField(r, "uint32", &m.PointStep); err != nil {
		return err
	}

	// RowStep
	if err = ros.DeserializeMessageField(r, "uint32", &m.RowStep); err != nil {
		return err
	}

	// Data
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Data: %s", err)
		}
		m.Data = make([]uint8, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "uint8", &m.Data[i]); err != nil {
				return err
			}
		}
	}

	// IsDense
	if err = ros.DeserializeMessageField(r, "bool", &m.IsDense); err != nil {
		return err
	}

	return
}
