package device

import (
	"bytes"
	"encoding/json"

	"github.com/na4ma4/go-hass-mqtt/model"
	"github.com/na4ma4/go-hass-mqtt/model/topic"
)

type Device struct {
	baseDeviceTopic *topic.Topic `json:"-"`
	hassTopic       *topic.Topic `json:"-"`

	ID              model.BasicIdentifier `json:"ids,omitempty"`
	Name            string                `json:"name,omitempty"`
	Manufacturer    string                `json:"mf,omitempty"`
	Model           string                `json:"mdl,omitempty"`
	SoftwareVersion string                `json:"sw,omitempty"`
	HardwareVersion string                `json:"hw,omitempty"`
	Serial          string                `json:"sn,omitempty"`
}

func New(id model.BasicIdentifier, opts ...OptFunc) *Device {
	d := &Device{
		ID:        id,
		Name:      id.String(),
		hassTopic: topic.New("homeassistant").Join("device", id.String()).Ptr(),
	}

	for _, opt := range opts {
		opt(d)
	}

	return d
}

func (d *Device) TopicHomeAssistant() topic.Topic {
	if d.hassTopic != nil {
		return *d.hassTopic
	}

	return topic.New("")
}

func (d *Device) TopicDeviceBase() topic.Topic {
	if d.baseDeviceTopic != nil {
		return *d.baseDeviceTopic
	}

	return topic.New("")
}

func (d *Device) ToByteBuffer() (*bytes.Buffer, error) {
	buf := bytes.NewBuffer(nil)
	err := json.NewEncoder(buf).Encode(d)
	return buf, err
}

func (d *Device) ToJSON() ([]byte, error) {
	buf, err := d.ToByteBuffer()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), err
}
