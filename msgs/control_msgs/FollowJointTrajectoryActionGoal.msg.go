// Code generated by ros-gen-go.
// source: FollowJointTrajectoryActionGoal.msg
// DO NOT EDIT!
package control_msgs

import (
	"io"

	"github.com/ppg/rosgo/ros"
	"github.com/ppg/rosgo/msgs/actionlib_msgs"
	"github.com/ppg/rosgo/msgs/std_msgs"
)

type _MsgFollowJointTrajectoryActionGoal struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgFollowJointTrajectoryActionGoal) Text() string {
	return t.text
}

func (t *_MsgFollowJointTrajectoryActionGoal) Name() string {
	return t.name
}

func (t *_MsgFollowJointTrajectoryActionGoal) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgFollowJointTrajectoryActionGoal) NewMessage() ros.Message {
	m := new(FollowJointTrajectoryActionGoal)

	return m
}

var (
	MsgFollowJointTrajectoryActionGoal = &_MsgFollowJointTrajectoryActionGoal{
		`# ====== DO NOT MODIFY! AUTOGENERATED FROM AN ACTION DEFINITION ======

Header header
actionlib_msgs/GoalID goal_id
FollowJointTrajectoryGoal goal
`,
		"control_msgs/FollowJointTrajectoryActionGoal",
		"ea97e347ff08322557a8ca22058939c7",
	}
)

type FollowJointTrajectoryActionGoal struct {
	Header std_msgs.Header
	GoalID actionlib_msgs.GoalID
	Goal   FollowJointTrajectoryGoal
}

func (m *FollowJointTrajectoryActionGoal) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "actionlib_msgs/GoalID", &m.GoalID); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "FollowJointTrajectoryGoal", &m.Goal); err != nil {
		return err
	}

	return
}

func (m *FollowJointTrajectoryActionGoal) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// GoalID
	if err = ros.DeserializeMessageField(r, "actionlib_msgs/GoalID", &m.GoalID); err != nil {
		return err
	}

	// Goal
	if err = ros.DeserializeMessageField(r, "FollowJointTrajectoryGoal", &m.Goal); err != nil {
		return err
	}

	return
}
