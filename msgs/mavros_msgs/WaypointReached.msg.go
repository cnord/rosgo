// Code generated by ros-gen-go.
// source: WaypointReached.msg
// DO NOT EDIT!
package mavros_msgs

import (
	"io"

	"github.com/cnord/rosgo/msgs/std_msgs"
	"github.com/cnord/rosgo/ros"
)

type _MsgWaypointReached struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgWaypointReached) Text() string {
	return t.text
}

func (t *_MsgWaypointReached) Name() string {
	return t.name
}

func (t *_MsgWaypointReached) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgWaypointReached) NewMessage() ros.Message {
	m := new(WaypointReached)

	return m
}

var (
	MsgWaypointReached = &_MsgWaypointReached{
		`# That message represent MISSION_ITEM_REACHED
#
#   :wp_seq:    index number of reached waypoint

Header header

uint16 wp_seq
`,
		"mavros_msgs/WaypointReached",
		"1cf64d072d9f6aa0a6715922bdde6a35",
	}
)

type WaypointReached struct {
	Header std_msgs.Header
	WpSeq  uint16
}

func (m *WaypointReached) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint16", &m.WpSeq); err != nil {
		return err
	}

	return
}

func (m *WaypointReached) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// WpSeq
	if err = ros.DeserializeMessageField(r, "uint16", &m.WpSeq); err != nil {
		return err
	}

	return
}
