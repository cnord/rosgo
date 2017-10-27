// Code generated by ros-gen-go.
// source: GlobalPositionTarget.msg
// DO NOT EDIT!
package mavros_msgs

import (
	"io"

	"github.com/cnord/rosgo/msgs/geometry_msgs"
	"github.com/cnord/rosgo/msgs/std_msgs"
	"github.com/cnord/rosgo/ros"
)

type _MsgGlobalPositionTarget struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgGlobalPositionTarget) Text() string {
	return t.text
}

func (t *_MsgGlobalPositionTarget) Name() string {
	return t.name
}

func (t *_MsgGlobalPositionTarget) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgGlobalPositionTarget) NewMessage() ros.Message {
	m := new(GlobalPositionTarget)

	return m
}

var (
	MsgGlobalPositionTarget = &_MsgGlobalPositionTarget{
		`# Message for SET_POSITION_TARGET_GLOBAL_INT
#
# Some complex system requires all feautures that mavlink
# message provide. See issue #402, #415.

std_msgs/Header header

uint8 coordinate_frame
uint8 FRAME_GLOBAL_INT = 5
uint8 FRAME_GLOBAL_REL_ALT = 6
uint8 FRAME_GLOBAL_TERRAIN_ALT = 11

uint16 type_mask
uint16 IGNORE_LATITUDE = 1	# Position ignore flags
uint16 IGNORE_LONGITUDE = 2
uint16 IGNORE_ALTITUDE = 4
uint16 IGNORE_VX = 8	# Velocity vector ignore flags
uint16 IGNORE_VY = 16
uint16 IGNORE_VZ = 32
uint16 IGNORE_AFX = 64	# Acceleration/Force vector ignore flags
uint16 IGNORE_AFY = 128
uint16 IGNORE_AFZ = 256
uint16 FORCE = 512	# Force in af vector flag
uint16 IGNORE_YAW = 1024
uint16 IGNORE_YAW_RATE = 2048

float64 latitude
float64 longitude
float32 altitude	# in meters, AMSL or above terrain
geometry_msgs/Vector3 velocity
geometry_msgs/Vector3 acceleration_or_force
float32 yaw
float32 yaw_rate
`,
		"mavros_msgs/GlobalPositionTarget",
		"076ded0190b9e045f9c55264659ef102",
	}
)

type GlobalPositionTarget struct {
	Header              std_msgs.Header
	CoordinateFrame     uint8
	TypeMask            uint16
	Latitude            float64
	Longitude           float64
	Altitude            float32
	Velocity            geometry_msgs.Vector3
	AccelerationOrForce geometry_msgs.Vector3
	Yaw                 float32
	YawRate             float32
}

func (m *GlobalPositionTarget) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "std_msgs/Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint8", &m.CoordinateFrame); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint16", &m.TypeMask); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float64", &m.Latitude); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float64", &m.Longitude); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.Altitude); err != nil {
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

func (m *GlobalPositionTarget) Deserialize(r io.Reader) (err error) {
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

	// Latitude
	if err = ros.DeserializeMessageField(r, "float64", &m.Latitude); err != nil {
		return err
	}

	// Longitude
	if err = ros.DeserializeMessageField(r, "float64", &m.Longitude); err != nil {
		return err
	}

	// Altitude
	if err = ros.DeserializeMessageField(r, "float32", &m.Altitude); err != nil {
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