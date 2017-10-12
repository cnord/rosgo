// Code generated by ros-gen-go.
// source: ProjectedMapInfo.msg
// DO NOT EDIT!
package map_msgs

import (
	"io"

	"github.com/cnord/rosgo/ros"
)

type _MsgProjectedMapInfo struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgProjectedMapInfo) Text() string {
	return t.text
}

func (t *_MsgProjectedMapInfo) Name() string {
	return t.name
}

func (t *_MsgProjectedMapInfo) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgProjectedMapInfo) NewMessage() ros.Message {
	m := new(ProjectedMapInfo)

	return m
}

var (
	MsgProjectedMapInfo = &_MsgProjectedMapInfo{
		`string frame_id
float64 x
float64 y
float64 width
float64 height
float64 min_z
float64 max_z`,
		"map_msgs/ProjectedMapInfo",
		"2dc10595ae94de23f22f8a6d2a0eef7a",
	}
)

type ProjectedMapInfo struct {
	FrameId string
	X       float64
	Y       float64
	Width   float64
	Height  float64
	MinZ    float64
	MaxZ    float64
}

func (m *ProjectedMapInfo) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "string", &m.FrameId); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float64", &m.X); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float64", &m.Y); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float64", &m.Width); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float64", &m.Height); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float64", &m.MinZ); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float64", &m.MaxZ); err != nil {
		return err
	}

	return
}

func (m *ProjectedMapInfo) Deserialize(r io.Reader) (err error) {
	// FrameId
	if err = ros.DeserializeMessageField(r, "string", &m.FrameId); err != nil {
		return err
	}

	// X
	if err = ros.DeserializeMessageField(r, "float64", &m.X); err != nil {
		return err
	}

	// Y
	if err = ros.DeserializeMessageField(r, "float64", &m.Y); err != nil {
		return err
	}

	// Width
	if err = ros.DeserializeMessageField(r, "float64", &m.Width); err != nil {
		return err
	}

	// Height
	if err = ros.DeserializeMessageField(r, "float64", &m.Height); err != nil {
		return err
	}

	// MinZ
	if err = ros.DeserializeMessageField(r, "float64", &m.MinZ); err != nil {
		return err
	}

	// MaxZ
	if err = ros.DeserializeMessageField(r, "float64", &m.MaxZ); err != nil {
		return err
	}

	return
}
