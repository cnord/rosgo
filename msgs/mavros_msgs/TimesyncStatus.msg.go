// Code generated by ros-gen-go.
// source: TimesyncStatus.msg
// DO NOT EDIT!
package mavros_msgs

import (
	"io"

	"github.com/cnord/rosgo/msgs/std_msgs"
	"github.com/cnord/rosgo/ros"
)

type _MsgTimesyncStatus struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgTimesyncStatus) Text() string {
	return t.text
}

func (t *_MsgTimesyncStatus) Name() string {
	return t.name
}

func (t *_MsgTimesyncStatus) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgTimesyncStatus) NewMessage() ros.Message {
	m := new(TimesyncStatus)

	return m
}

var (
	MsgTimesyncStatus = &_MsgTimesyncStatus{
		`# Status of the MAVLink time synchroniser

std_msgs/Header header
uint64 remote_timestamp_ns		# remote system timestamp (nanoseconds)
int64 observed_offset_ns		# raw time offset directly observed from this timesync packet (nanoseconds)
int64 estimated_offset_ns		# smoothed time offset between companion system and Mavros (nanoseconds)
float32 round_trip_time_ms		# round trip time of this timesync packet (milliseconds)`,
		"mavros_msgs/TimesyncStatus",
		"021ec8044e747bea518b441f374ba64b",
	}
)

type TimesyncStatus struct {
	Header            std_msgs.Header
	RemoteTimestampNs uint64
	ObservedOffsetNs  int64
	EstimatedOffsetNs int64
	RoundTripTimeMs   float32
}

func (m *TimesyncStatus) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "std_msgs/Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint64", &m.RemoteTimestampNs); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "int64", &m.ObservedOffsetNs); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "int64", &m.EstimatedOffsetNs); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.RoundTripTimeMs); err != nil {
		return err
	}

	return
}

func (m *TimesyncStatus) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "std_msgs/Header", &m.Header); err != nil {
		return err
	}

	// RemoteTimestampNs
	if err = ros.DeserializeMessageField(r, "uint64", &m.RemoteTimestampNs); err != nil {
		return err
	}

	// ObservedOffsetNs
	if err = ros.DeserializeMessageField(r, "int64", &m.ObservedOffsetNs); err != nil {
		return err
	}

	// EstimatedOffsetNs
	if err = ros.DeserializeMessageField(r, "int64", &m.EstimatedOffsetNs); err != nil {
		return err
	}

	// RoundTripTimeMs
	if err = ros.DeserializeMessageField(r, "float32", &m.RoundTripTimeMs); err != nil {
		return err
	}

	return
}
