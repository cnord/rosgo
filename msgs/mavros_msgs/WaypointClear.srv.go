// Code generated by ros-gen-go.
// source: WaypointClear.srv
// DO NOT EDIT!
package mavros_msgs

import (
	"io"

	"github.com/cnord/rosgo/ros"
)

// Service type metadata
type _SrvWaypointClear struct {
	name    string
	md5sum  string
	text    string
	reqType ros.MessageType
	resType ros.MessageType
}

func (t *_SrvWaypointClear) Name() string                  { return t.name }
func (t *_SrvWaypointClear) MD5Sum() string                { return t.md5sum }
func (t *_SrvWaypointClear) Text() string                  { return t.text }
func (t *_SrvWaypointClear) RequestType() ros.MessageType  { return t.reqType }
func (t *_SrvWaypointClear) ResponseType() ros.MessageType { return t.resType }
func (t *_SrvWaypointClear) NewService() ros.Service {
	return new(WaypointClear)
}

var (
	SrvWaypointClear = &_SrvWaypointClear{
		"mavros_msgs/WaypointClear",
		"358e233cde0c8a8bcfea4ce193f8fc15",
		`# Request clear waypoint

---
bool success
`,
		MsgWaypointClearRequest,
		MsgWaypointClearResponse,
	}
)

type WaypointClear struct {
	Request  WaypointClearRequest
	Response WaypointClearResponse
}

func (s *WaypointClear) ReqMessage() ros.Message { return &s.Request }
func (s *WaypointClear) ResMessage() ros.Message { return &s.Response }

// WaypointClearRequest

type _MsgWaypointClearRequest struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgWaypointClearRequest) Text() string {
	return t.text
}

func (t *_MsgWaypointClearRequest) Name() string {
	return t.name
}

func (t *_MsgWaypointClearRequest) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgWaypointClearRequest) NewMessage() ros.Message {
	m := new(WaypointClearRequest)

	return m
}

var (
	MsgWaypointClearRequest = &_MsgWaypointClearRequest{
		`# Request clear waypoint

`,
		"mavros_msgs/WaypointClearRequest",
		"358e233cde0c8a8bcfea4ce193f8fc15",
	}
)

type WaypointClearRequest struct {
}

func (m *WaypointClearRequest) Serialize(w io.Writer) (err error) {
	return
}

func (m *WaypointClearRequest) Deserialize(r io.Reader) (err error) {
	return
}

// WaypointClearResponse

type _MsgWaypointClearResponse struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgWaypointClearResponse) Text() string {
	return t.text
}

func (t *_MsgWaypointClearResponse) Name() string {
	return t.name
}

func (t *_MsgWaypointClearResponse) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgWaypointClearResponse) NewMessage() ros.Message {
	m := new(WaypointClearResponse)

	return m
}

var (
	MsgWaypointClearResponse = &_MsgWaypointClearResponse{
		`
bool success
`,
		"mavros_msgs/WaypointClearResponse",
		"358e233cde0c8a8bcfea4ce193f8fc15",
	}
)

type WaypointClearResponse struct {
	Success bool
}

func (m *WaypointClearResponse) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "bool", &m.Success); err != nil {
		return err
	}

	return
}

func (m *WaypointClearResponse) Deserialize(r io.Reader) (err error) {
	// Success
	if err = ros.DeserializeMessageField(r, "bool", &m.Success); err != nil {
		return err
	}

	return
}