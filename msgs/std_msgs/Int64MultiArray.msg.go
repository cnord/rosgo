// Code generated by ros-gen-go.
// source: Int64MultiArray.msg
// DO NOT EDIT!
package std_msgs

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/ppg/rosgo/ros"
)

type _MsgInt64MultiArray struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgInt64MultiArray) Text() string {
	return t.text
}

func (t *_MsgInt64MultiArray) Name() string {
	return t.name
}

func (t *_MsgInt64MultiArray) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgInt64MultiArray) NewMessage() ros.Message {
	m := new(Int64MultiArray)

	return m
}

var (
	MsgInt64MultiArray = &_MsgInt64MultiArray{
		`# Please look at the MultiArrayLayout message definition for
# documentation on all multiarrays.

MultiArrayLayout  layout        # specification of data layout
int64[]           data          # array of data

`,
		"std_msgs/Int64MultiArray",
		"3cf06fa72eb413c56790f4045031f8c9",
	}
)

type Int64MultiArray struct {
	Layout MultiArrayLayout
	Data   []int64
}

func (m *Int64MultiArray) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "MultiArrayLayout", &m.Layout); err != nil {
		return err
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Data)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Data {
		if err = ros.SerializeMessageField(w, "int64", &elem); err != nil {
			return err
		}
	}

	return
}

func (m *Int64MultiArray) Deserialize(r io.Reader) (err error) {
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
		m.Data = make([]int64, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "int64", &m.Data[i]); err != nil {
				return err
			}
		}
	}

	return
}