// Code generated by ros-gen-go.
// source: PointHeadFeedback.msg
// DO NOT EDIT!
package control_msgs

import (
	"io"

	"github.com/cnord/rosgo/ros"
)

type _MsgPointHeadFeedback struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgPointHeadFeedback) Text() string {
	return t.text
}

func (t *_MsgPointHeadFeedback) Name() string {
	return t.name
}

func (t *_MsgPointHeadFeedback) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgPointHeadFeedback) NewMessage() ros.Message {
	m := new(PointHeadFeedback)

	return m
}

var (
	MsgPointHeadFeedback = &_MsgPointHeadFeedback{
		`# ====== DO NOT MODIFY! AUTOGENERATED FROM AN ACTION DEFINITION ======
float64 pointing_angle_error

`,
		"control_msgs/PointHeadFeedback",
		"cce80d27fd763682da8805a73316cab4",
	}
)

type PointHeadFeedback struct {
	PointingAngleError float64
}

func (m *PointHeadFeedback) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "float64", &m.PointingAngleError); err != nil {
		return err
	}

	return
}

func (m *PointHeadFeedback) Deserialize(r io.Reader) (err error) {
	// PointingAngleError
	if err = ros.DeserializeMessageField(r, "float64", &m.PointingAngleError); err != nil {
		return err
	}

	return
}
