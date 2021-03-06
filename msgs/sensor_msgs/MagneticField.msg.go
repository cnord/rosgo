// Code generated by ros-gen-go.
// source: MagneticField.msg
// DO NOT EDIT!
package sensor_msgs

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/cnord/rosgo/msgs/geometry_msgs"
	"github.com/cnord/rosgo/msgs/std_msgs"
	"github.com/cnord/rosgo/ros"
)

type _MsgMagneticField struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgMagneticField) Text() string {
	return t.text
}

func (t *_MsgMagneticField) Name() string {
	return t.name
}

func (t *_MsgMagneticField) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgMagneticField) NewMessage() ros.Message {
	m := new(MagneticField)

	return m
}

var (
	MsgMagneticField = &_MsgMagneticField{
		` # Measurement of the Magnetic Field vector at a specific location.

 # If the covariance of the measurement is known, it should be filled in
 # (if all you know is the variance of each measurement, e.g. from the datasheet,
 #just put those along the diagonal)
 # A covariance matrix of all zeros will be interpreted as "covariance unknown",
 # and to use the data a covariance will have to be assumed or gotten from some
 # other source


 Header header                        # timestamp is the time the
                                      # field was measured
                                      # frame_id is the location and orientation
                                      # of the field measurement

 geometry_msgs/Vector3 magnetic_field # x, y, and z components of the
                                      # field vector in Tesla
                                      # If your sensor does not output 3 axes,
                                      # put NaNs in the components not reported.

 float64[9] magnetic_field_covariance # Row major about x, y, z axes
                                      # 0 is interpreted as variance unknown`,
		"sensor_msgs/MagneticField",
		"2f3b0b43eed0c9501de0fa3ff89a45aa",
	}
)

type MagneticField struct {
	Header                  std_msgs.Header
	MagneticField           geometry_msgs.Vector3
	MagneticFieldCovariance [9]float64
}

func (m *MagneticField) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "geometry_msgs/Vector3", &m.MagneticField); err != nil {
		return err
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.MagneticFieldCovariance)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.MagneticFieldCovariance {
		if err = ros.SerializeMessageField(w, "float64", &elem); err != nil {
			return err
		}
	}

	return
}

func (m *MagneticField) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// MagneticField
	if err = ros.DeserializeMessageField(r, "geometry_msgs/Vector3", &m.MagneticField); err != nil {
		return err
	}

	// MagneticFieldCovariance
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for MagneticFieldCovariance: %s", err)
		}
		if size > 9 {
			return fmt.Errorf("array size for MagneticFieldCovariance too large: expected=9, got=%d", size)
		}
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "float64", &m.MagneticFieldCovariance[i]); err != nil {
				return err
			}
		}
	}

	return
}
