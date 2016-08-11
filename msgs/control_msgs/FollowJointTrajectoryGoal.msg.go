// Code generated by ros-gen-go.
// source: FollowJointTrajectoryGoal.msg
// DO NOT EDIT!
package control_msgs

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/ppg/rosgo/msgs/trajectory_msgs"
	"github.com/ppg/rosgo/ros"
)

type _MsgFollowJointTrajectoryGoal struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgFollowJointTrajectoryGoal) Text() string {
	return t.text
}

func (t *_MsgFollowJointTrajectoryGoal) Name() string {
	return t.name
}

func (t *_MsgFollowJointTrajectoryGoal) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgFollowJointTrajectoryGoal) NewMessage() ros.Message {
	m := new(FollowJointTrajectoryGoal)

	return m
}

var (
	MsgFollowJointTrajectoryGoal = &_MsgFollowJointTrajectoryGoal{
		`# ====== DO NOT MODIFY! AUTOGENERATED FROM AN ACTION DEFINITION ======
# The joint trajectory to follow
trajectory_msgs/JointTrajectory trajectory

# Tolerances for the trajectory.  If the measured joint values fall
# outside the tolerances the trajectory goal is aborted.  Any
# tolerances that are not specified (by being omitted or set to 0) are
# set to the defaults for the action server (often taken from the
# parameter server).

# Tolerances applied to the joints as the trajectory is executed.  If
# violated, the goal aborts with error_code set to
# PATH_TOLERANCE_VIOLATED.
JointTolerance[] path_tolerance

# To report success, the joints must be within goal_tolerance of the
# final trajectory value.  The goal must be achieved by time the
# trajectory ends plus goal_time_tolerance.  (goal_time_tolerance
# allows some leeway in time, so that the trajectory goal can still
# succeed even if the joints reach the goal some time after the
# precise end time of the trajectory).
#
# If the joints are not within goal_tolerance after "trajectory finish
# time" + goal_time_tolerance, the goal aborts with error_code set to
# GOAL_TOLERANCE_VIOLATED
JointTolerance[] goal_tolerance
duration goal_time_tolerance

`,
		"control_msgs/FollowJointTrajectoryGoal",
		"d6e925ca987733b7452947e90e68dd68",
	}
)

type FollowJointTrajectoryGoal struct {
	Trajectory        trajectory_msgs.JointTrajectory
	PathTolerance     []JointTolerance
	GoalTolerance     []JointTolerance
	GoalTimeTolerance ros.Duration
}

func (m *FollowJointTrajectoryGoal) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "trajectory_msgs/JointTrajectory", &m.Trajectory); err != nil {
		return err
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.PathTolerance)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.PathTolerance {
		if err = ros.SerializeMessageField(w, "JointTolerance", &elem); err != nil {
			return err
		}
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.GoalTolerance)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.GoalTolerance {
		if err = ros.SerializeMessageField(w, "JointTolerance", &elem); err != nil {
			return err
		}
	}

	if err = ros.SerializeMessageField(w, "duration", &m.GoalTimeTolerance); err != nil {
		return err
	}

	return
}

func (m *FollowJointTrajectoryGoal) Deserialize(r io.Reader) (err error) {
	// Trajectory
	if err = ros.DeserializeMessageField(r, "trajectory_msgs/JointTrajectory", &m.Trajectory); err != nil {
		return err
	}

	// PathTolerance
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for PathTolerance: %s", err)
		}
		m.PathTolerance = make([]JointTolerance, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "JointTolerance", &m.PathTolerance[i]); err != nil {
				return err
			}
		}
	}

	// GoalTolerance
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for GoalTolerance: %s", err)
		}
		m.GoalTolerance = make([]JointTolerance, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "JointTolerance", &m.GoalTolerance[i]); err != nil {
				return err
			}
		}
	}

	// GoalTimeTolerance
	if err = ros.DeserializeMessageField(r, "duration", &m.GoalTimeTolerance); err != nil {
		return err
	}

	return
}
