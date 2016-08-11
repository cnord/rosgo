// Code generated by ros-gen-go.
// source: DiagnosticStatus.msg
// DO NOT EDIT!
package diagnostic_msgs

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/ppg/rosgo/ros"
)

type _MsgDiagnosticStatus struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgDiagnosticStatus) Text() string {
	return t.text
}

func (t *_MsgDiagnosticStatus) Name() string {
	return t.name
}

func (t *_MsgDiagnosticStatus) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgDiagnosticStatus) NewMessage() ros.Message {
	m := new(DiagnosticStatus)

	return m
}

var (
	MsgDiagnosticStatus = &_MsgDiagnosticStatus{
		`# This message holds the status of an individual component of the robot.
# 

# Possible levels of operations
byte OK=0
byte WARN=1
byte ERROR=2
byte STALE=3

byte level # level of operation enumerated above 
string name # a description of the test/component reporting
string message # a description of the status
string hardware_id # a hardware unique string
KeyValue[] values # an array of values associated with the status

`,
		"diagnostic_msgs/DiagnosticStatus",
		"bfaff5a20cf26a7d7255cbaa01dae095",
	}
)

type DiagnosticStatus struct {
	Level      int8
	Name       string
	Message    string
	HardwareID string
	Values     []KeyValue
}

func (m *DiagnosticStatus) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "byte", &m.Level); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "string", &m.Name); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "string", &m.Message); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "string", &m.HardwareID); err != nil {
		return err
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Values)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Values {
		if err = ros.SerializeMessageField(w, "KeyValue", &elem); err != nil {
			return err
		}
	}

	return
}

func (m *DiagnosticStatus) Deserialize(r io.Reader) (err error) {
	// Level
	if err = ros.DeserializeMessageField(r, "byte", &m.Level); err != nil {
		return err
	}

	// Name
	if err = ros.DeserializeMessageField(r, "string", &m.Name); err != nil {
		return err
	}

	// Message
	if err = ros.DeserializeMessageField(r, "string", &m.Message); err != nil {
		return err
	}

	// HardwareID
	if err = ros.DeserializeMessageField(r, "string", &m.HardwareID); err != nil {
		return err
	}

	// Values
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Values: %s", err)
		}
		m.Values = make([]KeyValue, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "KeyValue", &m.Values[i]); err != nil {
				return err
			}
		}
	}

	return
}
