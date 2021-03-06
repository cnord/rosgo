// Code generated by ros-gen-go.
// source: JointTrajectoryGoal.msg
// DO NOT EDIT!
package control_msgs

import (
	"io"

	"github.com/cnord/rosgo/msgs/trajectory_msgs"
	"github.com/cnord/rosgo/ros"
)

type _MsgJointTrajectoryGoal struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgJointTrajectoryGoal) Text() string {
	return t.text
}

func (t *_MsgJointTrajectoryGoal) Name() string {
	return t.name
}

func (t *_MsgJointTrajectoryGoal) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgJointTrajectoryGoal) NewMessage() ros.Message {
	m := new(JointTrajectoryGoal)

	return m
}

var (
	MsgJointTrajectoryGoal = &_MsgJointTrajectoryGoal{
		`# ====== DO NOT MODIFY! AUTOGENERATED FROM AN ACTION DEFINITION ======
trajectory_msgs/JointTrajectory trajectory
`,
		"control_msgs/JointTrajectoryGoal",
		"2a0eff76c870e8595636c2a562ca298e",
	}
)

type JointTrajectoryGoal struct {
	Trajectory trajectory_msgs.JointTrajectory
}

func (m *JointTrajectoryGoal) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "trajectory_msgs/JointTrajectory", &m.Trajectory); err != nil {
		return err
	}

	return
}

func (m *JointTrajectoryGoal) Deserialize(r io.Reader) (err error) {
	// Trajectory
	if err = ros.DeserializeMessageField(r, "trajectory_msgs/JointTrajectory", &m.Trajectory); err != nil {
		return err
	}

	return
}
