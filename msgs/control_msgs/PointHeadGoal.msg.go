// Code generated by ros-gen-go.
// source: PointHeadGoal.msg
// DO NOT EDIT!
package control_msgs

import (
	"io"

	"github.com/ppg/rosgo/msgs/geometry_msgs"
	"github.com/ppg/rosgo/ros"
)

type _MsgPointHeadGoal struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgPointHeadGoal) Text() string {
	return t.text
}

func (t *_MsgPointHeadGoal) Name() string {
	return t.name
}

func (t *_MsgPointHeadGoal) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgPointHeadGoal) NewMessage() ros.Message {
	m := new(PointHeadGoal)

	return m
}

var (
	MsgPointHeadGoal = &_MsgPointHeadGoal{
		`# ====== DO NOT MODIFY! AUTOGENERATED FROM AN ACTION DEFINITION ======
geometry_msgs/PointStamped target
geometry_msgs/Vector3 pointing_axis
string pointing_frame
duration min_duration
float64 max_velocity
`,
		"control_msgs/PointHeadGoal",
		"57514543f8e8c7f50339a60ad52c4524",
	}
)

type PointHeadGoal struct {
	Target        geometry_msgs.PointStamped
	PointingAxis  geometry_msgs.Vector3
	PointingFrame string
	MinDuration   ros.Duration
	MaxVelocity   float64
}

func (m *PointHeadGoal) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "geometry_msgs/PointStamped", &m.Target); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "geometry_msgs/Vector3", &m.PointingAxis); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "string", &m.PointingFrame); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "duration", &m.MinDuration); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float64", &m.MaxVelocity); err != nil {
		return err
	}

	return
}

func (m *PointHeadGoal) Deserialize(r io.Reader) (err error) {
	// Target
	if err = ros.DeserializeMessageField(r, "geometry_msgs/PointStamped", &m.Target); err != nil {
		return err
	}

	// PointingAxis
	if err = ros.DeserializeMessageField(r, "geometry_msgs/Vector3", &m.PointingAxis); err != nil {
		return err
	}

	// PointingFrame
	if err = ros.DeserializeMessageField(r, "string", &m.PointingFrame); err != nil {
		return err
	}

	// MinDuration
	if err = ros.DeserializeMessageField(r, "duration", &m.MinDuration); err != nil {
		return err
	}

	// MaxVelocity
	if err = ros.DeserializeMessageField(r, "float64", &m.MaxVelocity); err != nil {
		return err
	}

	return
}
