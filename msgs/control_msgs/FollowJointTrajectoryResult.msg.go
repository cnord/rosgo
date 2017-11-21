// Code generated by ros-gen-go.
// source: FollowJointTrajectoryResult.msg
// DO NOT EDIT!
package control_msgs

import (
	"io"

	"github.com/cnord/rosgo/ros"
)

const (
	SUCCESSFUL int32 = 0
)

type _MsgFollowJointTrajectoryResult struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgFollowJointTrajectoryResult) Text() string {
	return t.text
}

func (t *_MsgFollowJointTrajectoryResult) Name() string {
	return t.name
}

func (t *_MsgFollowJointTrajectoryResult) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgFollowJointTrajectoryResult) NewMessage() ros.Message {
	m := new(FollowJointTrajectoryResult)

	return m
}

var (
	MsgFollowJointTrajectoryResult = &_MsgFollowJointTrajectoryResult{
		`# ====== DO NOT MODIFY! AUTOGENERATED FROM AN ACTION DEFINITION ======
int32 error_code
int32 SUCCESSFUL = 0
int32 INVALID_GOAL = -1
int32 INVALID_JOINTS = -2
int32 OLD_HEADER_TIMESTAMP = -3
int32 PATH_TOLERANCE_VIOLATED = -4
int32 GOAL_TOLERANCE_VIOLATED = -5

# Human readable description of the error code. Contains complementary
# information that is especially useful when execution fails, for instance:
# - INVALID_GOAL: The reason for the invalid goal (e.g., the requested
#   trajectory is in the past).
# - INVALID_JOINTS: The mismatch between the expected controller joints
#   and those provided in the goal.
# - PATH_TOLERANCE_VIOLATED and GOAL_TOLERANCE_VIOLATED: Which joint
#   violated which tolerance, and by how much.
string error_string

`,
		"control_msgs/FollowJointTrajectoryResult",
		"493383b18409bfb604b4e26c676401d2",
	}
)

type FollowJointTrajectoryResult struct {
	ErrorCode             int32
	INVALIDGOAL           int32
	INVALIDJOINTS         int32
	OLDHEADERTIMESTAMP    int32
	PATHTOLERANCEVIOLATED int32
	GOALTOLERANCEVIOLATED int32
	ErrorString           string
}

func (m *FollowJointTrajectoryResult) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "int32", &m.ErrorCode); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "int32", &m.INVALIDGOAL); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "int32", &m.INVALIDJOINTS); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "int32", &m.OLDHEADERTIMESTAMP); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "int32", &m.PATHTOLERANCEVIOLATED); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "int32", &m.GOALTOLERANCEVIOLATED); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "string", &m.ErrorString); err != nil {
		return err
	}

	return
}

func (m *FollowJointTrajectoryResult) Deserialize(r io.Reader) (err error) {
	// ErrorCode
	if err = ros.DeserializeMessageField(r, "int32", &m.ErrorCode); err != nil {
		return err
	}

	// INVALIDGOAL
	if err = ros.DeserializeMessageField(r, "int32", &m.INVALIDGOAL); err != nil {
		return err
	}

	// INVALIDJOINTS
	if err = ros.DeserializeMessageField(r, "int32", &m.INVALIDJOINTS); err != nil {
		return err
	}

	// OLDHEADERTIMESTAMP
	if err = ros.DeserializeMessageField(r, "int32", &m.OLDHEADERTIMESTAMP); err != nil {
		return err
	}

	// PATHTOLERANCEVIOLATED
	if err = ros.DeserializeMessageField(r, "int32", &m.PATHTOLERANCEVIOLATED); err != nil {
		return err
	}

	// GOALTOLERANCEVIOLATED
	if err = ros.DeserializeMessageField(r, "int32", &m.GOALTOLERANCEVIOLATED); err != nil {
		return err
	}

	// ErrorString
	if err = ros.DeserializeMessageField(r, "string", &m.ErrorString); err != nil {
		return err
	}

	return
}
