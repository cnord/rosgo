// Code generated by ros-gen-go.
// source: ProjectedMap.msg
// DO NOT EDIT!
package map_msgs

import (
	"io"

	"github.com/ppg/rosgo/ros"
	"github.com/ppg/rosgo/msgs/nav_msgs"
)

type _MsgProjectedMap struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgProjectedMap) Text() string {
	return t.text
}

func (t *_MsgProjectedMap) Name() string {
	return t.name
}

func (t *_MsgProjectedMap) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgProjectedMap) NewMessage() ros.Message {
	m := new(ProjectedMap)

	return m
}

var (
	MsgProjectedMap = &_MsgProjectedMap{
		`nav_msgs/OccupancyGrid map
float64 min_z
float64 max_z`,
		"map_msgs/ProjectedMap",
		"39cfd3d65f74c1179ef7f7258f002167",
	}
)

type ProjectedMap struct {
	Map  nav_msgs.OccupancyGrid
	MinZ float64
	MaxZ float64
}

func (m *ProjectedMap) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "nav_msgs/OccupancyGrid", &m.Map); err != nil {
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

func (m *ProjectedMap) Deserialize(r io.Reader) (err error) {
	// Map
	if err = ros.DeserializeMessageField(r, "nav_msgs/OccupancyGrid", &m.Map); err != nil {
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
