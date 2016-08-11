// Code generated by ros-gen-go.
// source: InteractiveMarker.msg
// DO NOT EDIT!
package visualization_msgs

import (
	"io"
	"encoding/binary"
	"fmt"

	"github.com/ppg/rosgo/ros"
	"github.com/ppg/rosgo/msgs/geometry_msgs"
	"github.com/ppg/rosgo/msgs/std_msgs"
)

type _MsgInteractiveMarker struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgInteractiveMarker) Text() string {
	return t.text
}

func (t *_MsgInteractiveMarker) Name() string {
	return t.name
}

func (t *_MsgInteractiveMarker) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgInteractiveMarker) NewMessage() ros.Message {
	m := new(InteractiveMarker)

	return m
}

var (
	MsgInteractiveMarker = &_MsgInteractiveMarker{
		`# Time/frame info.
# If header.time is set to 0, the marker will be retransformed into
# its frame on each timestep. You will receive the pose feedback
# in the same frame.
# Otherwise, you might receive feedback in a different frame.
# For rviz, this will be the current 'fixed frame' set by the user.
Header header

# Initial pose. Also, defines the pivot point for rotations.
geometry_msgs/Pose pose

# Identifying string. Must be globally unique in
# the topic that this message is sent through.
string name

# Short description (< 40 characters).
string description

# Scale to be used for default controls (default=1).
float32 scale

# All menu and submenu entries associated with this marker.
MenuEntry[] menu_entries

# List of controls displayed for this marker.
InteractiveMarkerControl[] controls
`,
		"visualization_msgs/InteractiveMarker",
		"35474b6ee24280992e112232b62b4215",
	}
)

type InteractiveMarker struct {
	Header      std_msgs.Header
	Pose        geometry_msgs.Pose
	Name        string
	Description string
	Scale       float32
	MenuEntries []MenuEntry
	Controls    []InteractiveMarkerControl
}

func (m *InteractiveMarker) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "geometry_msgs/Pose", &m.Pose); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "string", &m.Name); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "string", &m.Description); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.Scale); err != nil {
		return err
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.MenuEntries)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.MenuEntries {
		if err = ros.SerializeMessageField(w, "MenuEntry", &elem); err != nil {
			return err
		}
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Controls)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Controls {
		if err = ros.SerializeMessageField(w, "InteractiveMarkerControl", &elem); err != nil {
			return err
		}
	}

	return
}

func (m *InteractiveMarker) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// Pose
	if err = ros.DeserializeMessageField(r, "geometry_msgs/Pose", &m.Pose); err != nil {
		return err
	}

	// Name
	if err = ros.DeserializeMessageField(r, "string", &m.Name); err != nil {
		return err
	}

	// Description
	if err = ros.DeserializeMessageField(r, "string", &m.Description); err != nil {
		return err
	}

	// Scale
	if err = ros.DeserializeMessageField(r, "float32", &m.Scale); err != nil {
		return err
	}

	// MenuEntries
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for MenuEntries: %s", err)
		}
		m.MenuEntries = make([]MenuEntry, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "MenuEntry", &m.MenuEntries[i]); err != nil {
				return err
			}
		}
	}

	// Controls
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Controls: %s", err)
		}
		m.Controls = make([]InteractiveMarkerControl, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "InteractiveMarkerControl", &m.Controls[i]); err != nil {
				return err
			}
		}
	}

	return
}
