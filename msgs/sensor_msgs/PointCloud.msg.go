// Code generated by ros-gen-go.
// source: PointCloud.msg
// DO NOT EDIT!
package sensor_msgs

import (
	"io"
	"encoding/binary"
	"fmt"

	"github.com/ppg/rosgo/ros"
	"github.com/ppg/rosgo/msgs/geometry_msgs"
	"github.com/ppg/rosgo/msgs/std_msgs"
)

type _MsgPointCloud struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgPointCloud) Text() string {
	return t.text
}

func (t *_MsgPointCloud) Name() string {
	return t.name
}

func (t *_MsgPointCloud) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgPointCloud) NewMessage() ros.Message {
	m := new(PointCloud)

	return m
}

var (
	MsgPointCloud = &_MsgPointCloud{
		`# This message holds a collection of 3d points, plus optional additional
# information about each point.

# Time of sensor data acquisition, coordinate frame ID.
Header header

# Array of 3d points. Each Point32 should be interpreted as a 3d point
# in the frame given in the header.
geometry_msgs/Point32[] points

# Each channel should have the same number of elements as points array,
# and the data in each channel should correspond 1:1 with each point.
# Channel names in common practice are listed in ChannelFloat32.msg.
ChannelFloat32[] channels
`,
		"sensor_msgs/PointCloud",
		"8fd83ffb348b4c6b507856d5fdaa54c2",
	}
)

type PointCloud struct {
	Header   std_msgs.Header
	Points   []geometry_msgs.Point32
	Channels []ChannelFloat32
}

func (m *PointCloud) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Points)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Points {
		if err = ros.SerializeMessageField(w, "geometry_msgs/Point32", &elem); err != nil {
			return err
		}
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Channels)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Channels {
		if err = ros.SerializeMessageField(w, "ChannelFloat32", &elem); err != nil {
			return err
		}
	}

	return
}

func (m *PointCloud) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// Points
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Points: %s", err)
		}
		m.Points = make([]geometry_msgs.Point32, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "geometry_msgs/Point32", &m.Points[i]); err != nil {
				return err
			}
		}
	}

	// Channels
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Channels: %s", err)
		}
		m.Channels = make([]ChannelFloat32, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "ChannelFloat32", &m.Channels[i]); err != nil {
				return err
			}
		}
	}

	return
}
