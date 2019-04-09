// Code generated by ros-gen-go.
// source: LogRequestData.srv
// DO NOT EDIT!
package mavros_msgs

import (
	"io"

	"github.com/cnord/rosgo/ros"
)

// Service type metadata
type _SrvLogRequestData struct {
	name    string
	md5sum  string
	text    string
	reqType ros.MessageType
	resType ros.MessageType
}

func (t *_SrvLogRequestData) Name() string                  { return t.name }
func (t *_SrvLogRequestData) MD5Sum() string                { return t.md5sum }
func (t *_SrvLogRequestData) Text() string                  { return t.text }
func (t *_SrvLogRequestData) RequestType() ros.MessageType  { return t.reqType }
func (t *_SrvLogRequestData) ResponseType() ros.MessageType { return t.resType }
func (t *_SrvLogRequestData) NewService() ros.Service {
	return new(LogRequestData)
}

var (
	SrvLogRequestData = &_SrvLogRequestData{
		"mavros_msgs/LogRequestData",
		"91170669bcd08438f0440f9b4d12a11a",
		`# Request a chunk of a log
#
#  :id: - log id from LogEntry message
#  :offset: - offset into the log
#  :count: - number of bytes to get

uint16 id
uint32 offset
uint32 count
---
bool success
`,
		MsgLogRequestDataRequest,
		MsgLogRequestDataResponse,
	}
)

type LogRequestData struct {
	Request  LogRequestDataRequest
	Response LogRequestDataResponse
}

func (s *LogRequestData) ReqMessage() ros.Message { return &s.Request }
func (s *LogRequestData) ResMessage() ros.Message { return &s.Response }

// LogRequestDataRequest

type _MsgLogRequestDataRequest struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgLogRequestDataRequest) Text() string {
	return t.text
}

func (t *_MsgLogRequestDataRequest) Name() string {
	return t.name
}

func (t *_MsgLogRequestDataRequest) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgLogRequestDataRequest) NewMessage() ros.Message {
	m := new(LogRequestDataRequest)

	return m
}

var (
	MsgLogRequestDataRequest = &_MsgLogRequestDataRequest{
		`# Request a chunk of a log
#
#  :id: - log id from LogEntry message
#  :offset: - offset into the log
#  :count: - number of bytes to get

uint16 id
uint32 offset
uint32 count
`,
		"mavros_msgs/LogRequestDataRequest",
		"",
	}
)

type LogRequestDataRequest struct {
	Id     uint16
	Offset uint32
	Count  uint32
}

func (m *LogRequestDataRequest) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "uint16", &m.Id); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint32", &m.Offset); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint32", &m.Count); err != nil {
		return err
	}

	return
}

func (m *LogRequestDataRequest) Deserialize(r io.Reader) (err error) {
	// Id
	if err = ros.DeserializeMessageField(r, "uint16", &m.Id); err != nil {
		return err
	}

	// Offset
	if err = ros.DeserializeMessageField(r, "uint32", &m.Offset); err != nil {
		return err
	}

	// Count
	if err = ros.DeserializeMessageField(r, "uint32", &m.Count); err != nil {
		return err
	}

	return
}

// LogRequestDataResponse

type _MsgLogRequestDataResponse struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgLogRequestDataResponse) Text() string {
	return t.text
}

func (t *_MsgLogRequestDataResponse) Name() string {
	return t.name
}

func (t *_MsgLogRequestDataResponse) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgLogRequestDataResponse) NewMessage() ros.Message {
	m := new(LogRequestDataResponse)

	return m
}

var (
	MsgLogRequestDataResponse = &_MsgLogRequestDataResponse{
		`
bool success
`,
		"mavros_msgs/LogRequestDataResponse",
		"",
	}
)

type LogRequestDataResponse struct {
	Success bool
}

func (m *LogRequestDataResponse) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "bool", &m.Success); err != nil {
		return err
	}

	return
}

func (m *LogRequestDataResponse) Deserialize(r io.Reader) (err error) {
	// Success
	if err = ros.DeserializeMessageField(r, "bool", &m.Success); err != nil {
		return err
	}

	return
}
