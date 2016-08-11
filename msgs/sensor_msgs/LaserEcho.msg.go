// Code generated by ros-gen-go.
// source: LaserEcho.msg
// DO NOT EDIT!
package sensor_msgs

import (
	"io"
	"encoding/binary"
	"fmt"

	"github.com/ppg/rosgo/ros"
)

type _MsgLaserEcho struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgLaserEcho) Text() string {
	return t.text
}

func (t *_MsgLaserEcho) Name() string {
	return t.name
}

func (t *_MsgLaserEcho) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgLaserEcho) NewMessage() ros.Message {
	m := new(LaserEcho)

	return m
}

var (
	MsgLaserEcho = &_MsgLaserEcho{
		`# This message is a submessage of MultiEchoLaserScan and is not intended
# to be used separately.

float32[] echoes  # Multiple values of ranges or intensities.
                  # Each array represents data from the same angle increment.`,
		"sensor_msgs/LaserEcho",
		"a3954a3a9a7d9a6a5e18f4d98439be09",
	}
)

type LaserEcho struct {
	Echoes []float32
}

func (m *LaserEcho) Serialize(w io.Writer) (err error) {
	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Echoes)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Echoes {
		if err = ros.SerializeMessageField(w, "float32", &elem); err != nil {
			return err
		}
	}

	return
}

func (m *LaserEcho) Deserialize(r io.Reader) (err error) {
	// Echoes
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Echoes: %s", err)
		}
		m.Echoes = make([]float32, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "float32", &m.Echoes[i]); err != nil {
				return err
			}
		}
	}

	return
}
