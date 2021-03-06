// Code generated by ros-gen-go.
// source: CommandTriggerInterval.srv
// DO NOT EDIT!
package mavros_msgs

import (
	"io"

	"github.com/cnord/rosgo/ros"
)

// Service type metadata
type _SrvCommandTriggerInterval struct {
	name    string
	md5sum  string
	text    string
	reqType ros.MessageType
	resType ros.MessageType
}

func (t *_SrvCommandTriggerInterval) Name() string                  { return t.name }
func (t *_SrvCommandTriggerInterval) MD5Sum() string                { return t.md5sum }
func (t *_SrvCommandTriggerInterval) Text() string                  { return t.text }
func (t *_SrvCommandTriggerInterval) RequestType() ros.MessageType  { return t.reqType }
func (t *_SrvCommandTriggerInterval) ResponseType() ros.MessageType { return t.resType }
func (t *_SrvCommandTriggerInterval) NewService() ros.Service {
	return new(CommandTriggerInterval)
}

var (
	SrvCommandTriggerInterval = &_SrvCommandTriggerInterval{
		"mavros_msgs/CommandTriggerInterval",
		"b16f28a04389d5d47ddaa9e025e7383a",
		`# Type for controlling camera trigger interval and integration time

float32   cycle_time			# Trigger cycle_time (interval between to triggers) - set to 0 to ignore command
float32   integration_time	# Camera shutter integration_time - set to 0 to ignore command
---
bool success
uint8 result`,
		MsgCommandTriggerIntervalRequest,
		MsgCommandTriggerIntervalResponse,
	}
)

type CommandTriggerInterval struct {
	Request  CommandTriggerIntervalRequest
	Response CommandTriggerIntervalResponse
}

func (s *CommandTriggerInterval) ReqMessage() ros.Message { return &s.Request }
func (s *CommandTriggerInterval) ResMessage() ros.Message { return &s.Response }

// CommandTriggerIntervalRequest

type _MsgCommandTriggerIntervalRequest struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgCommandTriggerIntervalRequest) Text() string {
	return t.text
}

func (t *_MsgCommandTriggerIntervalRequest) Name() string {
	return t.name
}

func (t *_MsgCommandTriggerIntervalRequest) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgCommandTriggerIntervalRequest) NewMessage() ros.Message {
	m := new(CommandTriggerIntervalRequest)

	return m
}

var (
	MsgCommandTriggerIntervalRequest = &_MsgCommandTriggerIntervalRequest{
		`# Type for controlling camera trigger interval and integration time

float32   cycle_time			# Trigger cycle_time (interval between to triggers) - set to 0 to ignore command
float32   integration_time	# Camera shutter integration_time - set to 0 to ignore command
`,
		"mavros_msgs/CommandTriggerIntervalRequest",
		"",
	}
)

type CommandTriggerIntervalRequest struct {
	CycleTime       float32
	IntegrationTime float32
}

func (m *CommandTriggerIntervalRequest) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "float32", &m.CycleTime); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.IntegrationTime); err != nil {
		return err
	}

	return
}

func (m *CommandTriggerIntervalRequest) Deserialize(r io.Reader) (err error) {
	// CycleTime
	if err = ros.DeserializeMessageField(r, "float32", &m.CycleTime); err != nil {
		return err
	}

	// IntegrationTime
	if err = ros.DeserializeMessageField(r, "float32", &m.IntegrationTime); err != nil {
		return err
	}

	return
}

// CommandTriggerIntervalResponse

type _MsgCommandTriggerIntervalResponse struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgCommandTriggerIntervalResponse) Text() string {
	return t.text
}

func (t *_MsgCommandTriggerIntervalResponse) Name() string {
	return t.name
}

func (t *_MsgCommandTriggerIntervalResponse) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgCommandTriggerIntervalResponse) NewMessage() ros.Message {
	m := new(CommandTriggerIntervalResponse)

	return m
}

var (
	MsgCommandTriggerIntervalResponse = &_MsgCommandTriggerIntervalResponse{
		`
bool success
uint8 result`,
		"mavros_msgs/CommandTriggerIntervalResponse",
		"",
	}
)

type CommandTriggerIntervalResponse struct {
	Success bool
	Result  uint8
}

func (m *CommandTriggerIntervalResponse) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "bool", &m.Success); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint8", &m.Result); err != nil {
		return err
	}

	return
}

func (m *CommandTriggerIntervalResponse) Deserialize(r io.Reader) (err error) {
	// Success
	if err = ros.DeserializeMessageField(r, "bool", &m.Success); err != nil {
		return err
	}

	// Result
	if err = ros.DeserializeMessageField(r, "uint8", &m.Result); err != nil {
		return err
	}

	return
}
