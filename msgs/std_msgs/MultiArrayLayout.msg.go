// Code generated by ros-gen-go.
// source: MultiArrayLayout.msg
// DO NOT EDIT!
package std_msgs

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/ppg/rosgo/ros"
)

type _MsgMultiArrayLayout struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgMultiArrayLayout) Text() string {
	return t.text
}

func (t *_MsgMultiArrayLayout) Name() string {
	return t.name
}

func (t *_MsgMultiArrayLayout) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgMultiArrayLayout) NewMessage() ros.Message {
	m := new(MultiArrayLayout)

	return m
}

var (
	MsgMultiArrayLayout = &_MsgMultiArrayLayout{
		`# The multiarray declares a generic multi-dimensional array of a
# particular data type.  Dimensions are ordered from outer most
# to inner most.

MultiArrayDimension[] dim # Array of dimension properties
uint32 data_offset        # padding elements at front of data

# Accessors should ALWAYS be written in terms of dimension stride
# and specified outer-most dimension first.
# 
# multiarray(i,j,k) = data[data_offset + dim_stride[1]*i + dim_stride[2]*j + k]
#
# A standard, 3-channel 640x480 image with interleaved color channels
# would be specified as:
#
# dim[0].label  = "height"
# dim[0].size   = 480
# dim[0].stride = 3*640*480 = 921600  (note dim[0] stride is just size of image)
# dim[1].label  = "width"
# dim[1].size   = 640
# dim[1].stride = 3*640 = 1920
# dim[2].label  = "channel"
# dim[2].size   = 3
# dim[2].stride = 3
#
# multiarray(i,j,k) refers to the ith row, jth column, and kth channel.
`,
		"std_msgs/MultiArrayLayout",
		"9e109a43d67cb5458e2e5239bb985dee",
	}
)

type MultiArrayLayout struct {
	Dim        []MultiArrayDimension
	DataOffset uint32
}

func (m *MultiArrayLayout) Serialize(w io.Writer) (err error) {
	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Dim)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Dim {
		if err = ros.SerializeMessageField(w, "MultiArrayDimension", &elem); err != nil {
			return err
		}
	}

	if err = ros.SerializeMessageField(w, "uint32", &m.DataOffset); err != nil {
		return err
	}

	return
}

func (m *MultiArrayLayout) Deserialize(r io.Reader) (err error) {
	// Dim
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Dim: %s", err)
		}
		m.Dim = make([]MultiArrayDimension, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "MultiArrayDimension", &m.Dim[i]); err != nil {
				return err
			}
		}
	}

	// DataOffset
	if err = ros.DeserializeMessageField(r, "uint32", &m.DataOffset); err != nil {
		return err
	}

	return
}
