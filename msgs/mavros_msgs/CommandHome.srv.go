// Code generated by ros-gen-go.
// source: CommandHome.srv
// DO NOT EDIT!
package mavros_msgs

import (
	"io"

	"github.com/cnord/rosgo/ros"
)

// Service type metadata
type _SrvCommandHome struct {
	name    string
	md5sum  string
	text    string
	reqType ros.MessageType
	resType ros.MessageType
}

func (t *_SrvCommandHome) Name() string                  { return t.name }
func (t *_SrvCommandHome) MD5Sum() string                { return t.md5sum }
func (t *_SrvCommandHome) Text() string                  { return t.text }
func (t *_SrvCommandHome) RequestType() ros.MessageType  { return t.reqType }
func (t *_SrvCommandHome) ResponseType() ros.MessageType { return t.resType }
func (t *_SrvCommandHome) NewService() ros.Service {
	return new(CommandHome)
}

var (
	SrvCommandHome = &_SrvCommandHome{
		"mavros_msgs/CommandHome",
		"3e89d38ecf947c9e62eb2e14dfa6e65d",
		`# request set new home position

bool current_gps
float32 latitude
float32 longitude
float32 altitude
---
bool success
uint8 result
`,
		MsgCommandHomeRequest,
		MsgCommandHomeResponse,
	}
)

type CommandHome struct {
	Request  CommandHomeRequest
	Response CommandHomeResponse
}

func (s *CommandHome) ReqMessage() ros.Message { return &s.Request }
func (s *CommandHome) ResMessage() ros.Message { return &s.Response }

// CommandHomeRequest

type _MsgCommandHomeRequest struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgCommandHomeRequest) Text() string {
	return t.text
}

func (t *_MsgCommandHomeRequest) Name() string {
	return t.name
}

func (t *_MsgCommandHomeRequest) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgCommandHomeRequest) NewMessage() ros.Message {
	m := new(CommandHomeRequest)

	return m
}

var (
	MsgCommandHomeRequest = &_MsgCommandHomeRequest{
		`# request set new home position

bool current_gps
float32 latitude
float32 longitude
float32 altitude
`,
		"mavros_msgs/CommandHomeRequest",
		"",
	}
)

type CommandHomeRequest struct {
	CurrentGps bool
	Latitude   float32
	Longitude  float32
	Altitude   float32
}

func (m *CommandHomeRequest) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "bool", &m.CurrentGps); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.Latitude); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.Longitude); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.Altitude); err != nil {
		return err
	}

	return
}

func (m *CommandHomeRequest) Deserialize(r io.Reader) (err error) {
	// CurrentGps
	if err = ros.DeserializeMessageField(r, "bool", &m.CurrentGps); err != nil {
		return err
	}

	// Latitude
	if err = ros.DeserializeMessageField(r, "float32", &m.Latitude); err != nil {
		return err
	}

	// Longitude
	if err = ros.DeserializeMessageField(r, "float32", &m.Longitude); err != nil {
		return err
	}

	// Altitude
	if err = ros.DeserializeMessageField(r, "float32", &m.Altitude); err != nil {
		return err
	}

	return
}

// CommandHomeResponse

type _MsgCommandHomeResponse struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgCommandHomeResponse) Text() string {
	return t.text
}

func (t *_MsgCommandHomeResponse) Name() string {
	return t.name
}

func (t *_MsgCommandHomeResponse) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgCommandHomeResponse) NewMessage() ros.Message {
	m := new(CommandHomeResponse)

	return m
}

var (
	MsgCommandHomeResponse = &_MsgCommandHomeResponse{
		`
bool success
uint8 result
`,
		"mavros_msgs/CommandHomeResponse",
		"",
	}
)

type CommandHomeResponse struct {
	Success bool
	Result  uint8
}

func (m *CommandHomeResponse) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "bool", &m.Success); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint8", &m.Result); err != nil {
		return err
	}

	return
}

func (m *CommandHomeResponse) Deserialize(r io.Reader) (err error) {
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
