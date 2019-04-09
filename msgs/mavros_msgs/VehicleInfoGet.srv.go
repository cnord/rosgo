// Code generated by ros-gen-go.
// source: VehicleInfoGet.srv
// DO NOT EDIT!
package mavros_msgs

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/cnord/rosgo/ros"
)

// Service type metadata
type _SrvVehicleInfoGet struct {
	name    string
	md5sum  string
	text    string
	reqType ros.MessageType
	resType ros.MessageType
}

func (t *_SrvVehicleInfoGet) Name() string                  { return t.name }
func (t *_SrvVehicleInfoGet) MD5Sum() string                { return t.md5sum }
func (t *_SrvVehicleInfoGet) Text() string                  { return t.text }
func (t *_SrvVehicleInfoGet) RequestType() ros.MessageType  { return t.reqType }
func (t *_SrvVehicleInfoGet) ResponseType() ros.MessageType { return t.resType }
func (t *_SrvVehicleInfoGet) NewService() ros.Service {
	return new(VehicleInfoGet)
}

var (
	SrvVehicleInfoGet = &_SrvVehicleInfoGet{
		"mavros_msgs/VehicleInfoGet",
		"519756d07eeab57c2f1ab9495e90e33f",
		`# Request the Vehicle Info
# use this to request the current target sysid / compid defined in mavros
# set get_all = True to request all available vehicles

uint8 GET_MY_SYSID = 0
uint8 GET_MY_COMPID = 0

uint8 sysid
uint8 compid
bool get_all
---
bool success
mavros_msgs/VehicleInfo[] vehicles

`,
		MsgVehicleInfoGetRequest,
		MsgVehicleInfoGetResponse,
	}
)

type VehicleInfoGet struct {
	Request  VehicleInfoGetRequest
	Response VehicleInfoGetResponse
}

func (s *VehicleInfoGet) ReqMessage() ros.Message { return &s.Request }
func (s *VehicleInfoGet) ResMessage() ros.Message { return &s.Response }

// VehicleInfoGetRequest

const (
	VehicleInfoGetRequest_GET_MY_SYSID  uint8 = 0
	VehicleInfoGetRequest_GET_MY_COMPID uint8 = 0
)

type _MsgVehicleInfoGetRequest struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgVehicleInfoGetRequest) Text() string {
	return t.text
}

func (t *_MsgVehicleInfoGetRequest) Name() string {
	return t.name
}

func (t *_MsgVehicleInfoGetRequest) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgVehicleInfoGetRequest) NewMessage() ros.Message {
	m := new(VehicleInfoGetRequest)

	return m
}

var (
	MsgVehicleInfoGetRequest = &_MsgVehicleInfoGetRequest{
		`# Request the Vehicle Info
# use this to request the current target sysid / compid defined in mavros
# set get_all = True to request all available vehicles

uint8 GET_MY_SYSID = 0
uint8 GET_MY_COMPID = 0

uint8 sysid
uint8 compid
bool get_all
`,
		"mavros_msgs/VehicleInfoGetRequest",
		"",
	}
)

type VehicleInfoGetRequest struct {
	Sysid  uint8
	Compid uint8
	GetAll bool
}

func (m *VehicleInfoGetRequest) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "uint8", &m.Sysid); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint8", &m.Compid); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "bool", &m.GetAll); err != nil {
		return err
	}

	return
}

func (m *VehicleInfoGetRequest) Deserialize(r io.Reader) (err error) {
	// Sysid
	if err = ros.DeserializeMessageField(r, "uint8", &m.Sysid); err != nil {
		return err
	}

	// Compid
	if err = ros.DeserializeMessageField(r, "uint8", &m.Compid); err != nil {
		return err
	}

	// GetAll
	if err = ros.DeserializeMessageField(r, "bool", &m.GetAll); err != nil {
		return err
	}

	return
}

// VehicleInfoGetResponse

type _MsgVehicleInfoGetResponse struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgVehicleInfoGetResponse) Text() string {
	return t.text
}

func (t *_MsgVehicleInfoGetResponse) Name() string {
	return t.name
}

func (t *_MsgVehicleInfoGetResponse) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgVehicleInfoGetResponse) NewMessage() ros.Message {
	m := new(VehicleInfoGetResponse)

	return m
}

var (
	MsgVehicleInfoGetResponse = &_MsgVehicleInfoGetResponse{
		`
bool success
mavros_msgs/VehicleInfo[] vehicles

`,
		"mavros_msgs/VehicleInfoGetResponse",
		"",
	}
)

type VehicleInfoGetResponse struct {
	Success  bool
	Vehicles []VehicleInfo
}

func (m *VehicleInfoGetResponse) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "bool", &m.Success); err != nil {
		return err
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Vehicles)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Vehicles {
		if err = ros.SerializeMessageField(w, "mavros_msgs/VehicleInfo", &elem); err != nil {
			return err
		}
	}

	return
}

func (m *VehicleInfoGetResponse) Deserialize(r io.Reader) (err error) {
	// Success
	if err = ros.DeserializeMessageField(r, "bool", &m.Success); err != nil {
		return err
	}

	// Vehicles
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Vehicles: %s", err)
		}
		m.Vehicles = make([]VehicleInfo, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "mavros_msgs/VehicleInfo", &m.Vehicles[i]); err != nil {
				return err
			}
		}
	}

	return
}