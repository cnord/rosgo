// Code generated by ros-gen-go.
// source: GetMapActionFeedback.msg
// DO NOT EDIT!
package nav_msgs

import (
	"io"

	"github.com/cnord/rosgo/msgs/actionlib_msgs"
	"github.com/cnord/rosgo/msgs/std_msgs"
	"github.com/cnord/rosgo/ros"
)

type _MsgGetMapActionFeedback struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgGetMapActionFeedback) Text() string {
	return t.text
}

func (t *_MsgGetMapActionFeedback) Name() string {
	return t.name
}

func (t *_MsgGetMapActionFeedback) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgGetMapActionFeedback) NewMessage() ros.Message {
	m := new(GetMapActionFeedback)

	return m
}

var (
	MsgGetMapActionFeedback = &_MsgGetMapActionFeedback{
		`# ====== DO NOT MODIFY! AUTOGENERATED FROM AN ACTION DEFINITION ======

Header header
actionlib_msgs/GoalStatus status
GetMapFeedback feedback
`,
		"nav_msgs/GetMapActionFeedback",
		"aae20e09065c3809e8a8e87c4c8953fd",
	}
)

type GetMapActionFeedback struct {
	Header   std_msgs.Header
	Status   actionlib_msgs.GoalStatus
	Feedback GetMapFeedback
}

func (m *GetMapActionFeedback) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "actionlib_msgs/GoalStatus", &m.Status); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "GetMapFeedback", &m.Feedback); err != nil {
		return err
	}

	return
}

func (m *GetMapActionFeedback) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// Status
	if err = ros.DeserializeMessageField(r, "actionlib_msgs/GoalStatus", &m.Status); err != nil {
		return err
	}

	// Feedback
	if err = ros.DeserializeMessageField(r, "GetMapFeedback", &m.Feedback); err != nil {
		return err
	}

	return
}
