// Code generated by ros-gen-go.
// source: Illuminance.msg
// DO NOT EDIT!
package sensor_msgs

import (
	"io"

	"github.com/ppg/rosgo/ros"
	"github.com/ppg/rosgo/msgs/std_msgs"
)

type _MsgIlluminance struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgIlluminance) Text() string {
	return t.text
}

func (t *_MsgIlluminance) Name() string {
	return t.name
}

func (t *_MsgIlluminance) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgIlluminance) NewMessage() ros.Message {
	m := new(Illuminance)

	return m
}

var (
	MsgIlluminance = &_MsgIlluminance{
		` # Single photometric illuminance measurement.  Light should be assumed to be
 # measured along the sensor's x-axis (the area of detection is the y-z plane).
 # The illuminance should have a 0 or positive value and be received with
 # the sensor's +X axis pointing toward the light source.

 # Photometric illuminance is the measure of the human eye's sensitivity of the
 # intensity of light encountering or passing through a surface.

 # All other Photometric and Radiometric measurements should
 # not use this message.
 # This message cannot represent:
 # Luminous intensity (candela/light source output)
 # Luminance (nits/light output per area)
 # Irradiance (watt/area), etc.

 Header header           # timestamp is the time the illuminance was measured
                         # frame_id is the location and direction of the reading

 float64 illuminance     # Measurement of the Photometric Illuminance in Lux.

 float64 variance        # 0 is interpreted as variance unknown`,
		"sensor_msgs/Illuminance",
		"10cf55f229d37dfb4f7e5cce8fb091ec",
	}
)

type Illuminance struct {
	Header      std_msgs.Header
	Illuminance float64
	Variance    float64
}

func (m *Illuminance) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float64", &m.Illuminance); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float64", &m.Variance); err != nil {
		return err
	}

	return
}

func (m *Illuminance) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// Illuminance
	if err = ros.DeserializeMessageField(r, "float64", &m.Illuminance); err != nil {
		return err
	}

	// Variance
	if err = ros.DeserializeMessageField(r, "float64", &m.Variance); err != nil {
		return err
	}

	return
}
