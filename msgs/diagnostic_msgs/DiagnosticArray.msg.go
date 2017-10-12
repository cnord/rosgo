// Code generated by ros-gen-go.
// source: DiagnosticArray.msg
// DO NOT EDIT!
package diagnostic_msgs

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/cnord/rosgo/msgs/std_msgs"
	"github.com/cnord/rosgo/ros"
)

type _MsgDiagnosticArray struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgDiagnosticArray) Text() string {
	return t.text
}

func (t *_MsgDiagnosticArray) Name() string {
	return t.name
}

func (t *_MsgDiagnosticArray) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgDiagnosticArray) NewMessage() ros.Message {
	m := new(DiagnosticArray)

	return m
}

var (
	MsgDiagnosticArray = &_MsgDiagnosticArray{
		`# This message is used to send diagnostic information about the state of the robot
Header header #for timestamp
DiagnosticStatus[] status # an array of components being reported on`,
		"diagnostic_msgs/DiagnosticArray",
		"60810da900de1dd6ddd437c3503511da",
	}
)

type DiagnosticArray struct {
	Header std_msgs.Header
	Status []DiagnosticStatus
}

func (m *DiagnosticArray) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Status)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Status {
		if err = ros.SerializeMessageField(w, "DiagnosticStatus", &elem); err != nil {
			return err
		}
	}

	return
}

func (m *DiagnosticArray) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// Status
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Status: %s", err)
		}
		m.Status = make([]DiagnosticStatus, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "DiagnosticStatus", &m.Status[i]); err != nil {
				return err
			}
		}
	}

	return
}
