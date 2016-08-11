// Code generated by ros-gen-go.
// source: JointTrajectoryPoint.msg
// DO NOT EDIT!
package trajectory_msgs

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/ppg/rosgo/ros"
)

type _MsgJointTrajectoryPoint struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgJointTrajectoryPoint) Text() string {
	return t.text
}

func (t *_MsgJointTrajectoryPoint) Name() string {
	return t.name
}

func (t *_MsgJointTrajectoryPoint) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgJointTrajectoryPoint) NewMessage() ros.Message {
	m := new(JointTrajectoryPoint)

	return m
}

var (
	MsgJointTrajectoryPoint = &_MsgJointTrajectoryPoint{
		`# Each trajectory point specifies either positions[, velocities[, accelerations]]
# or positions[, effort] for the trajectory to be executed.
# All specified values are in the same order as the joint names in JointTrajectory.msg

float64[] positions
float64[] velocities
float64[] accelerations
float64[] effort
duration time_from_start
`,
		"trajectory_msgs/JointTrajectoryPoint",
		"0b7b3d8dcc88390331fe8287246f673f",
	}
)

type JointTrajectoryPoint struct {
	Positions     []float64
	Velocities    []float64
	Accelerations []float64
	Effort        []float64
	TimeFromStart ros.Duration
}

func (m *JointTrajectoryPoint) Serialize(w io.Writer) (err error) {
	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Positions)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Positions {
		if err = ros.SerializeMessageField(w, "float64", &elem); err != nil {
			return err
		}
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Velocities)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Velocities {
		if err = ros.SerializeMessageField(w, "float64", &elem); err != nil {
			return err
		}
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Accelerations)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Accelerations {
		if err = ros.SerializeMessageField(w, "float64", &elem); err != nil {
			return err
		}
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Effort)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Effort {
		if err = ros.SerializeMessageField(w, "float64", &elem); err != nil {
			return err
		}
	}

	if err = ros.SerializeMessageField(w, "duration", &m.TimeFromStart); err != nil {
		return err
	}

	return
}

func (m *JointTrajectoryPoint) Deserialize(r io.Reader) (err error) {
	// Positions
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Positions: %s", err)
		}
		m.Positions = make([]float64, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "float64", &m.Positions[i]); err != nil {
				return err
			}
		}
	}

	// Velocities
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Velocities: %s", err)
		}
		m.Velocities = make([]float64, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "float64", &m.Velocities[i]); err != nil {
				return err
			}
		}
	}

	// Accelerations
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Accelerations: %s", err)
		}
		m.Accelerations = make([]float64, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "float64", &m.Accelerations[i]); err != nil {
				return err
			}
		}
	}

	// Effort
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Effort: %s", err)
		}
		m.Effort = make([]float64, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "float64", &m.Effort[i]); err != nil {
				return err
			}
		}
	}

	// TimeFromStart
	if err = ros.DeserializeMessageField(r, "duration", &m.TimeFromStart); err != nil {
		return err
	}

	return
}
