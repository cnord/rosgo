// Code generated by ros-gen-go.
// source: MultiArrayDimension.msg
// DO NOT EDIT!
package std_msgs

import (
	"io"

	"github.com/cnord/rosgo/ros"
)

type _MsgMultiArrayDimension struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgMultiArrayDimension) Text() string {
	return t.text
}

func (t *_MsgMultiArrayDimension) Name() string {
	return t.name
}

func (t *_MsgMultiArrayDimension) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgMultiArrayDimension) NewMessage() ros.Message {
	m := new(MultiArrayDimension)

	return m
}

var (
	MsgMultiArrayDimension = &_MsgMultiArrayDimension{
		`string label   # label of given dimension
uint32 size    # size of given dimension (in type units)
uint32 stride  # stride of given dimension`,
		"std_msgs/MultiArrayDimension",
		"4cd0c83a8683deae40ecdac60e53bfa8",
	}
)

type MultiArrayDimension struct {
	Label  string
	Size   uint32
	Stride uint32
}

func (m *MultiArrayDimension) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "string", &m.Label); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint32", &m.Size); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint32", &m.Stride); err != nil {
		return err
	}

	return
}

func (m *MultiArrayDimension) Deserialize(r io.Reader) (err error) {
	// Label
	if err = ros.DeserializeMessageField(r, "string", &m.Label); err != nil {
		return err
	}

	// Size
	if err = ros.DeserializeMessageField(r, "uint32", &m.Size); err != nil {
		return err
	}

	// Stride
	if err = ros.DeserializeMessageField(r, "uint32", &m.Stride); err != nil {
		return err
	}

	return
}
