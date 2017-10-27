// Code generated by ros-gen-go.
// source: PositionTarget.msg
// DO NOT EDIT!
package mavros_msgs

import (
	"io"

	"github.com/cnord/rosgo/msgs/geometry_msgs"
	"github.com/cnord/rosgo/msgs/std_msgs"
	"github.com/cnord/rosgo/ros"
)

type _MsgPositionTarget struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgPositionTarget) Text() string {
	return t.text
}

func (t *_MsgPositionTarget) Name() string {
	return t.name
}

func (t *_MsgPositionTarget) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgPositionTarget) NewMessage() ros.Message {
	m := new(PositionTarget)

	return m
}

var (
	MsgPositionTarget = &_MsgPositionTarget{
		`# Message for SET_POSITION_TARGET_LOCAL_NED
#
# Some complex system requires all feautures that mavlink
# message provide. See issue #402.

std_msgs/Header header

uint8 coordinate_frame
uint8 FRAME_LOCAL_NED = 1
uint8 FRAME_LOCAL_OFFSET_NED = 7
uint8 FRAME_BODY_NED = 8
uint8 FRAME_BODY_OFFSET_NED = 9

uint16 type_mask
uint16 IGNORE_PX = 1	# Position ignore flags
uint16 IGNORE_PY = 2
uint16 IGNORE_PZ = 4
uint16 IGNORE_VX = 8	# Velocity vector ignore flags
uint16 IGNORE_VY = 16
uint16 IGNORE_VZ = 32
uint16 IGNORE_AFX = 64	# Acceleration/Force vector ignore flags
uint16 IGNORE_AFY = 128
uint16 IGNORE_AFZ = 256
uint16 FORCE = 512	# Force in af vector flag
uint16 IGNORE_YAW = 1024
uint16 IGNORE_YAW_RATE = 2048

geometry_msgs/Point position
geometry_msgs/Vector3 velocity
geometry_msgs/Vector3 acceleration_or_force
float32 yaw
float32 yaw_rate
`,
		"mavros_msgs/PositionTarget",
		"dedb0081aaf8cb20209737746bb49117",
	}
)

type PositionTarget struct {
	Header              std_msgs.Header
	CoordinateFrame     uint8
	TypeMask            uint16
	Position            geometry_msgs.Point
	Velocity            geometry_msgs.Vector3
	AccelerationOrForce geometry_msgs.Vector3
	Yaw                 float32
	YawRate             float32
}

func (m *PositionTarget) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "std_msgs/Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint8", &m.CoordinateFrame); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint16", &m.TypeMask); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "geometry_msgs/Point", &m.Position); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "geometry_msgs/Vector3", &m.Velocity); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "geometry_msgs/Vector3", &m.AccelerationOrForce); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.Yaw); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.YawRate); err != nil {
		return err
	}

	return
}

func (m *PositionTarget) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "std_msgs/Header", &m.Header); err != nil {
		return err
	}

	// CoordinateFrame
	if err = ros.DeserializeMessageField(r, "uint8", &m.CoordinateFrame); err != nil {
		return err
	}

	// TypeMask
	if err = ros.DeserializeMessageField(r, "uint16", &m.TypeMask); err != nil {
		return err
	}

	// Position
	if err = ros.DeserializeMessageField(r, "geometry_msgs/Point", &m.Position); err != nil {
		return err
	}

	// Velocity
	if err = ros.DeserializeMessageField(r, "geometry_msgs/Vector3", &m.Velocity); err != nil {
		return err
	}

	// AccelerationOrForce
	if err = ros.DeserializeMessageField(r, "geometry_msgs/Vector3", &m.AccelerationOrForce); err != nil {
		return err
	}

	// Yaw
	if err = ros.DeserializeMessageField(r, "float32", &m.Yaw); err != nil {
		return err
	}

	// YawRate
	if err = ros.DeserializeMessageField(r, "float32", &m.YawRate); err != nil {
		return err
	}

	return
}