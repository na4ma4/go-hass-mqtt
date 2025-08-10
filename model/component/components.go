package component

import (
	"github.com/na4ma4/go-hass-mqtt/model"
	"github.com/na4ma4/go-hass-mqtt/model/topic"
)

type Component struct {
	BaseTopic *topic.Topic `json:"~,omitempty"`

	ID                  *model.BasicIdentifier `json:"uniq_id,omitempty"`
	Name                *string                `json:"name,omitempty"`
	Platform            *string                `json:"p,omitempty"`
	Mode                *string                `json:"mode,omitempty"`
	DeviceClass         *string                `json:"dev_cla,omitempty"`
	EnabledByDefault    *bool                  `json:"en,omitempty"`
	UnitOfMeasurement   *string                `json:"unit_of_meas,omitempty"`
	CommandTemplate     *string                `json:"cmd_tpl,omitempty"`
	ValueTemplate       *string                `json:"val_tpl,omitempty"`
	CommandTopic        *topic.Topic           `json:"cmd_t,omitempty"`
	StateTopic          *topic.Topic           `json:"stat_t,omitempty"`
	AvailabilityTopic   *topic.Topic           `json:"avty_t,omitempty"`
	Options             []string               `json:"ops,omitempty"`
	Maximum             *int64                 `json:"max,omitempty"`
	Minimum             *int64                 `json:"min,omitempty"`
	PayloadAvailable    *string                `json:"pl_avail,omitempty"`
	PayloadNotAvailable *string                `json:"pl_not_avail,omitempty"`
}

func New(id *model.BasicIdentifier, opts ...OptFunc) *Component {
	c := &Component{
		ID: id,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

type OptFunc func(*Component)

func WithCommandTemplate(template string) OptFunc {
	return func(c *Component) {
		c.CommandTemplate = &template
	}
}

func WithEnabledByDefault(enabledByDefault bool) OptFunc {
	return func(c *Component) {
		c.EnabledByDefault = &enabledByDefault
	}
}

func WithOptions(opts []string) OptFunc {
	return func(c *Component) {
		c.Options = opts
	}
}

func WithValueTemplate(template string) OptFunc {
	return func(c *Component) {
		c.ValueTemplate = &template
	}
}

func WithName(name string) OptFunc {
	return func(c *Component) {
		c.Name = &name
	}
}

func WithBaseTopic(topic topic.Topic) OptFunc {
	return func(c *Component) {
		c.BaseTopic = &topic
	}
}

func WithCommandTopic(topic topic.Topic) OptFunc {
	return func(c *Component) {
		c.CommandTopic = &topic
	}
}

func WithStateTopic(topic topic.Topic) OptFunc {
	return func(c *Component) {
		c.StateTopic = &topic
	}
}

func WithAvailabilityTopic(topic topic.Topic) OptFunc {
	return func(c *Component) {
		c.AvailabilityTopic = &topic
	}
}

func WithMinimum(min int64) OptFunc {
	return func(c *Component) {
		c.Minimum = &min
	}
}

func WithMaximum(max int64) OptFunc {
	return func(c *Component) {
		c.Maximum = &max
	}
}

func WithPayloadAvailable(payload string) OptFunc {
	return func(c *Component) {
		c.PayloadAvailable = &payload
	}
}

func WithPayloadNotAvailable(payload string) OptFunc {
	return func(c *Component) {
		c.PayloadNotAvailable = &payload
	}
}
