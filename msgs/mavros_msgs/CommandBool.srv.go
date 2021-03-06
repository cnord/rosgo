// Code generated by ros-gen-go.
// source: CommandBool.srv
// DO NOT EDIT!
package mavros_msgs

import (
	"io"

	"github.com/cnord/rosgo/ros"
)

// Service type metadata
type _SrvCommandBool struct {
	name    string
	md5sum  string
	text    string
	reqType ros.MessageType
	resType ros.MessageType
}

func (t *_SrvCommandBool) Name() string                  { return t.name }
func (t *_SrvCommandBool) MD5Sum() string                { return t.md5sum }
func (t *_SrvCommandBool) Text() string                  { return t.text }
func (t *_SrvCommandBool) RequestType() ros.MessageType  { return t.reqType }
func (t *_SrvCommandBool) ResponseType() ros.MessageType { return t.resType }
func (t *_SrvCommandBool) NewService() ros.Service {
	return new(CommandBool)
}

var (
	SrvCommandBool = &_SrvCommandBool{
		"mavros_msgs/CommandBool",
		"e09abbb4e5bae6b558e5010966eb6e9e",
		`# Common type for switch commands

bool value
---
bool success
uint8 result
`,
		MsgCommandBoolRequest,
		MsgCommandBoolResponse,
	}
)

type CommandBool struct {
	Request  CommandBoolRequest
	Response CommandBoolResponse
}

func (s *CommandBool) ReqMessage() ros.Message { return &s.Request }
func (s *CommandBool) ResMessage() ros.Message { return &s.Response }

// CommandBoolRequest

type _MsgCommandBoolRequest struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgCommandBoolRequest) Text() string {
	return t.text
}

func (t *_MsgCommandBoolRequest) Name() string {
	return t.name
}

func (t *_MsgCommandBoolRequest) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgCommandBoolRequest) NewMessage() ros.Message {
	m := new(CommandBoolRequest)

	return m
}

var (
	MsgCommandBoolRequest = &_MsgCommandBoolRequest{
		`# Common type for switch commands

bool value
`,
		"mavros_msgs/CommandBoolRequest",
		"",
	}
)

type CommandBoolRequest struct {
	Value bool
}

func (m *CommandBoolRequest) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "bool", &m.Value); err != nil {
		return err
	}

	return
}

func (m *CommandBoolRequest) Deserialize(r io.Reader) (err error) {
	// Value
	if err = ros.DeserializeMessageField(r, "bool", &m.Value); err != nil {
		return err
	}

	return
}

// CommandBoolResponse

type _MsgCommandBoolResponse struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgCommandBoolResponse) Text() string {
	return t.text
}

func (t *_MsgCommandBoolResponse) Name() string {
	return t.name
}

func (t *_MsgCommandBoolResponse) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgCommandBoolResponse) NewMessage() ros.Message {
	m := new(CommandBoolResponse)

	return m
}

var (
	MsgCommandBoolResponse = &_MsgCommandBoolResponse{
		`
bool success
uint8 result
`,
		"mavros_msgs/CommandBoolResponse",
		"",
	}
)

type CommandBoolResponse struct {
	Success bool
	Result  uint8
}

func (m *CommandBoolResponse) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "bool", &m.Success); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint8", &m.Result); err != nil {
		return err
	}

	return
}

func (m *CommandBoolResponse) Deserialize(r io.Reader) (err error) {
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
