// Code generated by ros-gen-go.
// source: TwistWithCovarianceStamped.msg
// DO NOT EDIT!
package geometry_msgs

import (
	"io"

	"github.com/ppg/rosgo/ros"
	"github.com/ppg/rosgo/msgs/std_msgs"
)

type _MsgTwistWithCovarianceStamped struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgTwistWithCovarianceStamped) Text() string {
	return t.text
}

func (t *_MsgTwistWithCovarianceStamped) Name() string {
	return t.name
}

func (t *_MsgTwistWithCovarianceStamped) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgTwistWithCovarianceStamped) NewMessage() ros.Message {
	m := new(TwistWithCovarianceStamped)

	return m
}

var (
	MsgTwistWithCovarianceStamped = &_MsgTwistWithCovarianceStamped{
		`# This represents an estimated twist with reference coordinate frame and timestamp.
Header header
TwistWithCovariance twist
`,
		"geometry_msgs/TwistWithCovarianceStamped",
		"f59b87b044187f26ef66329003c3d275",
	}
)

type TwistWithCovarianceStamped struct {
	Header std_msgs.Header
	Twist  TwistWithCovariance
}

func (m *TwistWithCovarianceStamped) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "TwistWithCovariance", &m.Twist); err != nil {
		return err
	}

	return
}

func (m *TwistWithCovarianceStamped) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// Twist
	if err = ros.DeserializeMessageField(r, "TwistWithCovariance", &m.Twist); err != nil {
		return err
	}

	return
}
