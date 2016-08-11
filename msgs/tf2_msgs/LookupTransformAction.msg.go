// Code generated by ros-gen-go.
// source: LookupTransformAction.msg
// DO NOT EDIT!
package tf2_msgs

import (
	"io"

	"github.com/ppg/rosgo/ros"
)

type _MsgLookupTransformAction struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgLookupTransformAction) Text() string {
	return t.text
}

func (t *_MsgLookupTransformAction) Name() string {
	return t.name
}

func (t *_MsgLookupTransformAction) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgLookupTransformAction) NewMessage() ros.Message {
	m := new(LookupTransformAction)

	return m
}

var (
	MsgLookupTransformAction = &_MsgLookupTransformAction{
		`# ====== DO NOT MODIFY! AUTOGENERATED FROM AN ACTION DEFINITION ======

LookupTransformActionGoal action_goal
LookupTransformActionResult action_result
LookupTransformActionFeedback action_feedback
`,
		"tf2_msgs/LookupTransformAction",
		"ea0b675693e01bf415f75c692ea50172",
	}
)

type LookupTransformAction struct {
	ActionGoal     LookupTransformActionGoal
	ActionResult   LookupTransformActionResult
	ActionFeedback LookupTransformActionFeedback
}

func (m *LookupTransformAction) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "LookupTransformActionGoal", &m.ActionGoal); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "LookupTransformActionResult", &m.ActionResult); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "LookupTransformActionFeedback", &m.ActionFeedback); err != nil {
		return err
	}

	return
}

func (m *LookupTransformAction) Deserialize(r io.Reader) (err error) {
	// ActionGoal
	if err = ros.DeserializeMessageField(r, "LookupTransformActionGoal", &m.ActionGoal); err != nil {
		return err
	}

	// ActionResult
	if err = ros.DeserializeMessageField(r, "LookupTransformActionResult", &m.ActionResult); err != nil {
		return err
	}

	// ActionFeedback
	if err = ros.DeserializeMessageField(r, "LookupTransformActionFeedback", &m.ActionFeedback); err != nil {
		return err
	}

	return
}
