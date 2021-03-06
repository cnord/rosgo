// Code generated by ros-gen-go.
// source: FileClose.srv
// DO NOT EDIT!
package mavros_msgs

import (
	"io"

	"github.com/cnord/rosgo/ros"
)

// Service type metadata
type _SrvFileClose struct {
	name    string
	md5sum  string
	text    string
	reqType ros.MessageType
	resType ros.MessageType
}

func (t *_SrvFileClose) Name() string                  { return t.name }
func (t *_SrvFileClose) MD5Sum() string                { return t.md5sum }
func (t *_SrvFileClose) Text() string                  { return t.text }
func (t *_SrvFileClose) RequestType() ros.MessageType  { return t.reqType }
func (t *_SrvFileClose) ResponseType() ros.MessageType { return t.resType }
func (t *_SrvFileClose) NewService() ros.Service {
	return new(FileClose)
}

var (
	SrvFileClose = &_SrvFileClose{
		"mavros_msgs/FileClose",
		"1c309002d7f7c6df7d3f094925ea8e10",
		`# FTP::Close
#
# Call FTP::Open first.
# :success:	indicates success end of request
# :r_errno:	remote errno if applicapable

string file_path
---
bool success
int32 r_errno
`,
		MsgFileCloseRequest,
		MsgFileCloseResponse,
	}
)

type FileClose struct {
	Request  FileCloseRequest
	Response FileCloseResponse
}

func (s *FileClose) ReqMessage() ros.Message { return &s.Request }
func (s *FileClose) ResMessage() ros.Message { return &s.Response }

// FileCloseRequest

type _MsgFileCloseRequest struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgFileCloseRequest) Text() string {
	return t.text
}

func (t *_MsgFileCloseRequest) Name() string {
	return t.name
}

func (t *_MsgFileCloseRequest) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgFileCloseRequest) NewMessage() ros.Message {
	m := new(FileCloseRequest)

	return m
}

var (
	MsgFileCloseRequest = &_MsgFileCloseRequest{
		`# FTP::Close
#
# Call FTP::Open first.
# :success:	indicates success end of request
# :r_errno:	remote errno if applicapable

string file_path
`,
		"mavros_msgs/FileCloseRequest",
		"",
	}
)

type FileCloseRequest struct {
	FilePath string
}

func (m *FileCloseRequest) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "string", &m.FilePath); err != nil {
		return err
	}

	return
}

func (m *FileCloseRequest) Deserialize(r io.Reader) (err error) {
	// FilePath
	if err = ros.DeserializeMessageField(r, "string", &m.FilePath); err != nil {
		return err
	}

	return
}

// FileCloseResponse

type _MsgFileCloseResponse struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgFileCloseResponse) Text() string {
	return t.text
}

func (t *_MsgFileCloseResponse) Name() string {
	return t.name
}

func (t *_MsgFileCloseResponse) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgFileCloseResponse) NewMessage() ros.Message {
	m := new(FileCloseResponse)

	return m
}

var (
	MsgFileCloseResponse = &_MsgFileCloseResponse{
		`
bool success
int32 r_errno
`,
		"mavros_msgs/FileCloseResponse",
		"",
	}
)

type FileCloseResponse struct {
	Success bool
	RErrno  int32
}

func (m *FileCloseResponse) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "bool", &m.Success); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "int32", &m.RErrno); err != nil {
		return err
	}

	return
}

func (m *FileCloseResponse) Deserialize(r io.Reader) (err error) {
	// Success
	if err = ros.DeserializeMessageField(r, "bool", &m.Success); err != nil {
		return err
	}

	// RErrno
	if err = ros.DeserializeMessageField(r, "int32", &m.RErrno); err != nil {
		return err
	}

	return
}
