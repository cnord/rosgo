// Code generated by ros-gen-go.
// source: InteractiveMarkerUpdate.msg
// DO NOT EDIT!
package visualization_msgs

import (
	"io"
	"encoding/binary"
	"fmt"

	"github.com/ppg/rosgo/ros"
)

type _MsgInteractiveMarkerUpdate struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgInteractiveMarkerUpdate) Text() string {
	return t.text
}

func (t *_MsgInteractiveMarkerUpdate) Name() string {
	return t.name
}

func (t *_MsgInteractiveMarkerUpdate) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgInteractiveMarkerUpdate) NewMessage() ros.Message {
	m := new(InteractiveMarkerUpdate)

	return m
}

var (
	MsgInteractiveMarkerUpdate = &_MsgInteractiveMarkerUpdate{
		`# Identifying string. Must be unique in the topic namespace
# that this server works on.
string server_id

# Sequence number.
# The client will use this to detect if it has missed an update.
uint64 seq_num

# Type holds the purpose of this message.  It must be one of UPDATE or KEEP_ALIVE.
# UPDATE: Incremental update to previous state. 
#         The sequence number must be 1 higher than for
#         the previous update.
# KEEP_ALIVE: Indicates the that the server is still living.
#             The sequence number does not increase.
#             No payload data should be filled out (markers, poses, or erases).
uint8 KEEP_ALIVE = 0
uint8 UPDATE = 1

uint8 type

#Note: No guarantees on the order of processing.
#      Contents must be kept consistent by sender.

#Markers to be added or updated
InteractiveMarker[] markers

#Poses of markers that should be moved
InteractiveMarkerPose[] poses

#Names of markers to be erased
string[] erases
`,
		"visualization_msgs/InteractiveMarkerUpdate",
		"78e02717d45a794bd09a18d1e530a0fe",
	}
)

type InteractiveMarkerUpdate struct {
	ServerID string
	SeqNum   uint64
	Type     uint8
	Markers  []InteractiveMarker
	Poses    []InteractiveMarkerPose
	Erases   []string
}

func (m *InteractiveMarkerUpdate) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "string", &m.ServerID); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint64", &m.SeqNum); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint8", &m.Type); err != nil {
		return err
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Markers)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Markers {
		if err = ros.SerializeMessageField(w, "InteractiveMarker", &elem); err != nil {
			return err
		}
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Poses)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Poses {
		if err = ros.SerializeMessageField(w, "InteractiveMarkerPose", &elem); err != nil {
			return err
		}
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Erases)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Erases {
		if err = ros.SerializeMessageField(w, "string", &elem); err != nil {
			return err
		}
	}

	return
}

func (m *InteractiveMarkerUpdate) Deserialize(r io.Reader) (err error) {
	// ServerID
	if err = ros.DeserializeMessageField(r, "string", &m.ServerID); err != nil {
		return err
	}

	// SeqNum
	if err = ros.DeserializeMessageField(r, "uint64", &m.SeqNum); err != nil {
		return err
	}

	// Type
	if err = ros.DeserializeMessageField(r, "uint8", &m.Type); err != nil {
		return err
	}

	// Markers
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Markers: %s", err)
		}
		m.Markers = make([]InteractiveMarker, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "InteractiveMarker", &m.Markers[i]); err != nil {
				return err
			}
		}
	}

	// Poses
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Poses: %s", err)
		}
		m.Poses = make([]InteractiveMarkerPose, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "InteractiveMarkerPose", &m.Poses[i]); err != nil {
				return err
			}
		}
	}

	// Erases
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Erases: %s", err)
		}
		m.Erases = make([]string, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "string", &m.Erases[i]); err != nil {
				return err
			}
		}
	}

	return
}
