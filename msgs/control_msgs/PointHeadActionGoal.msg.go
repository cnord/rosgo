// Code generated by ros-gen-go.
// source: PointHeadActionGoal.msg
// DO NOT EDIT!
package control_msgs

import (
	"io"

	"github.com/cnord/rosgo/msgs/actionlib_msgs"
	"github.com/cnord/rosgo/msgs/std_msgs"
	"github.com/cnord/rosgo/ros"
)

type _MsgPointHeadActionGoal struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgPointHeadActionGoal) Text() string {
	return t.text
}

func (t *_MsgPointHeadActionGoal) Name() string {
	return t.name
}

func (t *_MsgPointHeadActionGoal) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgPointHeadActionGoal) NewMessage() ros.Message {
	m := new(PointHeadActionGoal)

	return m
}

var (
	MsgPointHeadActionGoal = &_MsgPointHeadActionGoal{
		`# ====== DO NOT MODIFY! AUTOGENERATED FROM AN ACTION DEFINITION ======

Header header
actionlib_msgs/GoalID goal_id
PointHeadGoal goal
`,
		"control_msgs/PointHeadActionGoal",
		"b53a8323d0ba7b310ba17a2d3a82a6b8",
	}
)

type PointHeadActionGoal struct {
	Header std_msgs.Header
	GoalId actionlib_msgs.GoalID
	Goal   PointHeadGoal
}

func (m *PointHeadActionGoal) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "actionlib_msgs/GoalID", &m.GoalId); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "PointHeadGoal", &m.Goal); err != nil {
		return err
	}

	return
}

func (m *PointHeadActionGoal) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// GoalId
	if err = ros.DeserializeMessageField(r, "actionlib_msgs/GoalID", &m.GoalId); err != nil {
		return err
	}

	// Goal
	if err = ros.DeserializeMessageField(r, "PointHeadGoal", &m.Goal); err != nil {
		return err
	}

	return
}
