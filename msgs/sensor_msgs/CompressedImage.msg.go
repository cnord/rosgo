// Code generated by ros-gen-go.
// source: CompressedImage.msg
// DO NOT EDIT!
package sensor_msgs

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/cnord/rosgo/msgs/std_msgs"
	"github.com/cnord/rosgo/ros"
)

type _MsgCompressedImage struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgCompressedImage) Text() string {
	return t.text
}

func (t *_MsgCompressedImage) Name() string {
	return t.name
}

func (t *_MsgCompressedImage) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgCompressedImage) NewMessage() ros.Message {
	m := new(CompressedImage)

	return m
}

var (
	MsgCompressedImage = &_MsgCompressedImage{
		`# This message contains a compressed image

Header header        # Header timestamp should be acquisition time of image
                     # Header frame_id should be optical frame of camera
                     # origin of frame should be optical center of camera
                     # +x should point to the right in the image
                     # +y should point down in the image
                     # +z should point into to plane of the image

string format        # Specifies the format of the data
                     #   Acceptable values:
                     #     jpeg, png
uint8[] data         # Compressed image buffer
`,
		"sensor_msgs/CompressedImage",
		"8f7a12909da2c9d3332d540a0977563f",
	}
)

type CompressedImage struct {
	Header std_msgs.Header
	Format string
	Data   []uint8
}

func (m *CompressedImage) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "string", &m.Format); err != nil {
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

	return
}

func (m *CompressedImage) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// Format
	if err = ros.DeserializeMessageField(r, "string", &m.Format); err != nil {
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

	return
}
