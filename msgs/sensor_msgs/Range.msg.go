// Code generated by ros-gen-go.
// source: Range.msg
// DO NOT EDIT!
package sensor_msgs

import (
	"io"

	"github.com/ppg/rosgo/ros"
	"github.com/ppg/rosgo/msgs/std_msgs"
)

type _MsgRange struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgRange) Text() string {
	return t.text
}

func (t *_MsgRange) Name() string {
	return t.name
}

func (t *_MsgRange) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgRange) NewMessage() ros.Message {
	m := new(Range)

	return m
}

var (
	MsgRange = &_MsgRange{
		`# Single range reading from an active ranger that emits energy and reports
# one range reading that is valid along an arc at the distance measured. 
# This message is  not appropriate for laser scanners. See the LaserScan
# message if you are working with a laser scanner.

# This message also can represent a fixed-distance (binary) ranger.  This
# sensor will have min_range===max_range===distance of detection.
# These sensors follow REP 117 and will output -Inf if the object is detected
# and +Inf if the object is outside of the detection range.

Header header           # timestamp in the header is the time the ranger
                        # returned the distance reading

# Radiation type enums
# If you want a value added to this list, send an email to the ros-users list
uint8 ULTRASOUND=0
uint8 INFRARED=1

uint8 radiation_type    # the type of radiation used by the sensor
                        # (sound, IR, etc) [enum]

float32 field_of_view   # the size of the arc that the distance reading is
                        # valid for [rad]
                        # the object causing the range reading may have
                        # been anywhere within -field_of_view/2 and
                        # field_of_view/2 at the measured range. 
                        # 0 angle corresponds to the x-axis of the sensor.

float32 min_range       # minimum range value [m]
float32 max_range       # maximum range value [m]
                        # Fixed distance rangers require min_range==max_range

float32 range           # range data [m]
                        # (Note: values < range_min or > range_max
                        # should be discarded)
                        # Fixed distance rangers only output -Inf or +Inf.
                        # -Inf represents a detection within fixed distance.
                        # (Detection too close to the sensor to quantify)
                        # +Inf represents no detection within the fixed distance.
                        # (Object out of range)`,
		"sensor_msgs/Range",
		"71e981252e448b5283f03f85f1f485dc",
	}
)

type Range struct {
	Header        std_msgs.Header
	RadiationType uint8
	FieldOfView   float32
	MinRange      float32
	MaxRange      float32
	Range         float32
}

func (m *Range) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint8", &m.RadiationType); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.FieldOfView); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.MinRange); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.MaxRange); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.Range); err != nil {
		return err
	}

	return
}

func (m *Range) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// RadiationType
	if err = ros.DeserializeMessageField(r, "uint8", &m.RadiationType); err != nil {
		return err
	}

	// FieldOfView
	if err = ros.DeserializeMessageField(r, "float32", &m.FieldOfView); err != nil {
		return err
	}

	// MinRange
	if err = ros.DeserializeMessageField(r, "float32", &m.MinRange); err != nil {
		return err
	}

	// MaxRange
	if err = ros.DeserializeMessageField(r, "float32", &m.MaxRange); err != nil {
		return err
	}

	// Range
	if err = ros.DeserializeMessageField(r, "float32", &m.Range); err != nil {
		return err
	}

	return
}
