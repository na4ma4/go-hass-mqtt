package messages

import (
	"bytes"
	"encoding/json"

	"github.com/na4ma4/go-hass-mqtt/model/component"
	"github.com/na4ma4/go-hass-mqtt/model/device"
	"github.com/na4ma4/go-hass-mqtt/model/origin"
)

type Discovery struct {
	Device       *device.Device                  `json:"dev,omitempty"`
	Origin       *origin.Origin                  `json:"o,omitempty"`
	Cmps         map[string]*component.Component `json:"cmps,omitempty"`
	CommandTopic string                          `json:"cmd_t,omitempty"`
	StateTopic   string                          `json:"stat_t,omitempty"`
	Qos          byte                            `json:"qos,omitempty"`
}

func NewDiscovery(dev *device.Device, org *origin.Origin, opts ...DiscoveryOptFunc) *Discovery {
	d := &Discovery{
		// HomeTopic:    dev.TopicHome(),
		Device: dev,
		Origin: org,
		// StateTopic:   "~/state",
		// CommandTopic: "~/cmd",
		// StateTopic:   dev.TopicHome() + "/state",
		// CommandTopic: dev.TopicHome() + "/cmd",
	}

	for _, opt := range opts {
		opt(d)
	}

	return d
}

type DiscoveryOptFunc func(*Discovery)

func WithComponents(cmps ...*component.Component) DiscoveryOptFunc {
	return func(d *Discovery) {
		if d.Cmps == nil {
			d.Cmps = make(map[string]*component.Component)
		}

		for _, cmp := range cmps {
			if cmp.ID == nil {
				cmp.ID = &d.Device.ID
			}
			d.Cmps[cmp.ID.String()] = cmp
		}
	}
}

func WithQos(qos byte) DiscoveryOptFunc {
	return func(d *Discovery) {
		d.Qos = qos
	}
}

func WithCommandTopic(topic string) DiscoveryOptFunc {
	return func(d *Discovery) {
		d.CommandTopic = topic
	}
}

func WithStateTopic(topic string) DiscoveryOptFunc {
	return func(d *Discovery) {
		d.StateTopic = topic
	}
}

func (d *Discovery) Bytes() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(d); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (d *Discovery) BytesBuffer() (*bytes.Buffer, error) {
	buf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(d); err != nil {
		return nil, err
	}
	return buf, nil
}

// type Device struct {
// 	Ids  string `json:"ids"`
// 	Name string `json:"name"`
// 	Mf   string `json:"mf"`
// 	Mdl  string `json:"mdl"`
// 	Sw   string `json:"sw"`
// 	Sn   string `json:"sn"`
// 	Hw   string `json:"hw"`
// }

// type Origin struct {
// 	Name string `json:"name"`
// 	Sw   string `json:"sw"`
// 	URL  string `json:"url"`
// }

// type Component struct {
// 	Platform          *string `json:"p"`
// 	DeviceClass       *string `json:"device_class"`
// 	UnitOfMeasurement *string `json:"unit_of_measurement"`
// 	ValueTemplate     *string `json:"value_template"`
// 	UniqueID          *string `json:"unique_id"`
// }
