// Code generated by ros-gen-go.
// source: Int16MultiArray.msg
// DO NOT EDIT!
package std_msgs

import (
	"io"
	"encoding/binary"
	"fmt"

	"github.com/ppg/rosgo/ros"
)

type _MsgInt16MultiArray struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgInt16MultiArray) Text() string {
	return t.text
}

func (t *_MsgInt16MultiArray) Name() string {
	return t.name
}

func (t *_MsgInt16MultiArray) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgInt16MultiArray) NewMessage() ros.Message {
	m := new(Int16MultiArray)

	return m
}

var (
	MsgInt16MultiArray = &_MsgInt16MultiArray{
		`# Please look at the MultiArrayLayout message definition for
# documentation on all multiarrays.

MultiArrayLayout  layout        # specification of data layout
int16[]           data          # array of data

`,
		"std_msgs/Int16MultiArray",
		"1ea95620317e8e06fc1614473147a513",
	}
)

type Int16MultiArray struct {
	Layout MultiArrayLayout
	Data   []int16
}

func (m *Int16MultiArray) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "MultiArrayLayout", &m.Layout); err != nil {
		return err
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Data)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Data {
		if err = ros.SerializeMessageField(w, "int16", &elem); err != nil {
			return err
		}
	}

	return
}

func (m *Int16MultiArray) Deserialize(r io.Reader) (err error) {
	// Layout
	if err = ros.DeserializeMessageField(r, "MultiArrayLayout", &m.Layout); err != nil {
		return err
	}

	// Data
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Data: %s", err)
		}
		m.Data = make([]int16, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "int16", &m.Data[i]); err != nil {
				return err
			}
		}
	}

	return
}
