// Code generated by ros-gen-go.
// source: GripperCommandActionResult.msg
// DO NOT EDIT!
package control_msgs

import (
	"io"

	"github.com/cnord/rosgo/msgs/actionlib_msgs"
	"github.com/cnord/rosgo/msgs/std_msgs"
	"github.com/cnord/rosgo/ros"
)

type _MsgGripperCommandActionResult struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgGripperCommandActionResult) Text() string {
	return t.text
}

func (t *_MsgGripperCommandActionResult) Name() string {
	return t.name
}

func (t *_MsgGripperCommandActionResult) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgGripperCommandActionResult) NewMessage() ros.Message {
	m := new(GripperCommandActionResult)

	return m
}

var (
	MsgGripperCommandActionResult = &_MsgGripperCommandActionResult{
		`# ====== DO NOT MODIFY! AUTOGENERATED FROM AN ACTION DEFINITION ======

Header header
actionlib_msgs/GoalStatus status
GripperCommandResult result
`,
		"control_msgs/GripperCommandActionResult",
		"143702cb2df0f163c5283cedc5efc6b6",
	}
)

type GripperCommandActionResult struct {
	Header std_msgs.Header
	Status actionlib_msgs.GoalStatus
	Result GripperCommandResult
}

func (m *GripperCommandActionResult) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "actionlib_msgs/GoalStatus", &m.Status); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "GripperCommandResult", &m.Result); err != nil {
		return err
	}

	return
}

func (m *GripperCommandActionResult) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// Status
	if err = ros.DeserializeMessageField(r, "actionlib_msgs/GoalStatus", &m.Status); err != nil {
		return err
	}

	// Result
	if err = ros.DeserializeMessageField(r, "GripperCommandResult", &m.Result); err != nil {
		return err
	}

	return
}
