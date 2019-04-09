// Code generated by ros-gen-go.
// source: Pose2D.msg
// DO NOT EDIT!
package geometry_msgs

import (
	"io"

	"github.com/cnord/rosgo/ros"
)

type _MsgPose2D struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgPose2D) Text() string {
	return t.text
}

func (t *_MsgPose2D) Name() string {
	return t.name
}

func (t *_MsgPose2D) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgPose2D) NewMessage() ros.Message {
	m := new(Pose2D)

	return m
}

var (
	MsgPose2D = &_MsgPose2D{
		`# Deprecated
# Please use the full 3D pose.

# In general our recommendation is to use a full 3D representation of everything and for 2D specific applications make the appropriate projections into the plane for their calculations but optimally will preserve the 3D information during processing.

# If we have parallel copies of 2D datatypes every UI and other pipeline will end up needing to have dual interfaces to plot everything. And you will end up with not being able to use 3D tools for 2D use cases even if they're completely valid, as you'd have to reimplement it with different inputs and outputs. It's not particularly hard to plot the 2D pose or compute the yaw error for the Pose message and there are already tools and libraries that can do this for you.


# This expresses a position and orientation on a 2D manifold.

float64 x
float64 y
float64 theta
`,
		"geometry_msgs/Pose2D",
		"938fa65709584ad8e77d238529be13b8",
	}
)

type Pose2D struct {
	X     float64
	Y     float64
	Theta float64
}

func (m *Pose2D) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "float64", &m.X); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float64", &m.Y); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float64", &m.Theta); err != nil {
		return err
	}

	return
}

func (m *Pose2D) Deserialize(r io.Reader) (err error) {
	// X
	if err = ros.DeserializeMessageField(r, "float64", &m.X); err != nil {
		return err
	}

	// Y
	if err = ros.DeserializeMessageField(r, "float64", &m.Y); err != nil {
		return err
	}

	// Theta
	if err = ros.DeserializeMessageField(r, "float64", &m.Theta); err != nil {
		return err
	}

	return
}
