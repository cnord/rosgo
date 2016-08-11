// Code generated by ros-gen-go.
// source: GripperCommandActionGoal.msg
// DO NOT EDIT!
package control_msgs

import (
	"io"

	"github.com/ppg/rosgo/msgs/actionlib_msgs"
	"github.com/ppg/rosgo/msgs/std_msgs"
	"github.com/ppg/rosgo/ros"
)

type _MsgGripperCommandActionGoal struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgGripperCommandActionGoal) Text() string {
	return t.text
}

func (t *_MsgGripperCommandActionGoal) Name() string {
	return t.name
}

func (t *_MsgGripperCommandActionGoal) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgGripperCommandActionGoal) NewMessage() ros.Message {
	m := new(GripperCommandActionGoal)

	return m
}

var (
	MsgGripperCommandActionGoal = &_MsgGripperCommandActionGoal{
		`# ====== DO NOT MODIFY! AUTOGENERATED FROM AN ACTION DEFINITION ======

Header header
actionlib_msgs/GoalID goal_id
GripperCommandGoal goal
`,
		"control_msgs/GripperCommandActionGoal",
		"afbb0281892695d5519cbad8f2cc3fce",
	}
)

type GripperCommandActionGoal struct {
	Header std_msgs.Header
	GoalID actionlib_msgs.GoalID
	Goal   GripperCommandGoal
}

func (m *GripperCommandActionGoal) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "actionlib_msgs/GoalID", &m.GoalID); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "GripperCommandGoal", &m.Goal); err != nil {
		return err
	}

	return
}

func (m *GripperCommandActionGoal) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// GoalID
	if err = ros.DeserializeMessageField(r, "actionlib_msgs/GoalID", &m.GoalID); err != nil {
		return err
	}

	// Goal
	if err = ros.DeserializeMessageField(r, "GripperCommandGoal", &m.Goal); err != nil {
		return err
	}

	return
}
