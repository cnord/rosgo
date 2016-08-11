// Code generated by ros-gen-go.
// source: TF2Error.msg
// DO NOT EDIT!
package tf2_msgs

import (
	"io"

	"github.com/ppg/rosgo/ros"
)

type _MsgTF2Error struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgTF2Error) Text() string {
	return t.text
}

func (t *_MsgTF2Error) Name() string {
	return t.name
}

func (t *_MsgTF2Error) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgTF2Error) NewMessage() ros.Message {
	m := new(TF2Error)

	return m
}

var (
	MsgTF2Error = &_MsgTF2Error{
		`uint8 NO_ERROR = 0
uint8 LOOKUP_ERROR = 1
uint8 CONNECTIVITY_ERROR = 2
uint8 EXTRAPOLATION_ERROR = 3
uint8 INVALID_ARGUMENT_ERROR = 4
uint8 TIMEOUT_ERROR = 5
uint8 TRANSFORM_ERROR = 6

uint8 error
string error_string
`,
		"tf2_msgs/TF2Error",
		"4373ee33104a6a2611a3bce8d34370e7",
	}
)

type TF2Error struct {
	Error       uint8
	ErrorString string
}

func (m *TF2Error) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "uint8", &m.Error); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "string", &m.ErrorString); err != nil {
		return err
	}

	return
}

func (m *TF2Error) Deserialize(r io.Reader) (err error) {
	// Error
	if err = ros.DeserializeMessageField(r, "uint8", &m.Error); err != nil {
		return err
	}

	// ErrorString
	if err = ros.DeserializeMessageField(r, "string", &m.ErrorString); err != nil {
		return err
	}

	return
}
