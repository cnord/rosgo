// Code generated by ros-gen-go.
// source: HilSensor.msg
// DO NOT EDIT!
package mavros_msgs

import (
	"io"

	"github.com/cnord/rosgo/msgs/geometry_msgs"
	"github.com/cnord/rosgo/msgs/std_msgs"
	"github.com/cnord/rosgo/ros"
)

type _MsgHilSensor struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgHilSensor) Text() string {
	return t.text
}

func (t *_MsgHilSensor) Name() string {
	return t.name
}

func (t *_MsgHilSensor) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgHilSensor) NewMessage() ros.Message {
	m := new(HilSensor)

	return m
}

var (
	MsgHilSensor = &_MsgHilSensor{
		`# HilSensor.msg
#
# ROS representation of MAVLink HIL_SENSOR
# See mavlink message documentation here:
# https://mavlink.io/en/messages/common.html#HIL_SENSOR

std_msgs/Header header

geometry_msgs/Vector3 acc
geometry_msgs/Vector3 gyro
geometry_msgs/Vector3 mag
float32 abs_pressure
float32 diff_pressure
float32 pressure_alt
float32 temperature
uint32 fields_updated
`,
		"mavros_msgs/HilSensor",
		"2a892891e5c40d6dd1066bf1f394b5dc",
	}
)

type HilSensor struct {
	Header        std_msgs.Header
	Acc           geometry_msgs.Vector3
	Gyro          geometry_msgs.Vector3
	Mag           geometry_msgs.Vector3
	AbsPressure   float32
	DiffPressure  float32
	PressureAlt   float32
	Temperature   float32
	FieldsUpdated uint32
}

func (m *HilSensor) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "std_msgs/Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "geometry_msgs/Vector3", &m.Acc); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "geometry_msgs/Vector3", &m.Gyro); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "geometry_msgs/Vector3", &m.Mag); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.AbsPressure); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.DiffPressure); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.PressureAlt); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.Temperature); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint32", &m.FieldsUpdated); err != nil {
		return err
	}

	return
}

func (m *HilSensor) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "std_msgs/Header", &m.Header); err != nil {
		return err
	}

	// Acc
	if err = ros.DeserializeMessageField(r, "geometry_msgs/Vector3", &m.Acc); err != nil {
		return err
	}

	// Gyro
	if err = ros.DeserializeMessageField(r, "geometry_msgs/Vector3", &m.Gyro); err != nil {
		return err
	}

	// Mag
	if err = ros.DeserializeMessageField(r, "geometry_msgs/Vector3", &m.Mag); err != nil {
		return err
	}

	// AbsPressure
	if err = ros.DeserializeMessageField(r, "float32", &m.AbsPressure); err != nil {
		return err
	}

	// DiffPressure
	if err = ros.DeserializeMessageField(r, "float32", &m.DiffPressure); err != nil {
		return err
	}

	// PressureAlt
	if err = ros.DeserializeMessageField(r, "float32", &m.PressureAlt); err != nil {
		return err
	}

	// Temperature
	if err = ros.DeserializeMessageField(r, "float32", &m.Temperature); err != nil {
		return err
	}

	// FieldsUpdated
	if err = ros.DeserializeMessageField(r, "uint32", &m.FieldsUpdated); err != nil {
		return err
	}

	return
}
