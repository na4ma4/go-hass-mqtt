package mqttconn

import (
	"context"
	"net/url"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/na4ma4/go-hass-mqtt/model"
	"github.com/na4ma4/go-hass-mqtt/model/component"
	"github.com/na4ma4/go-hass-mqtt/model/device"
	"github.com/na4ma4/go-hass-mqtt/model/origin"
	"github.com/na4ma4/go-hass-mqtt/model/topic"
)

type (
	OnConnectHandler        func()
	OnConnectionLostHandler func(error)
	ClientOptions           struct {
		Servers           []*url.URL
		ClientID          string
		Username          string
		Password          string
		CleanSession      bool
		QoS               byte // Quality of Service level for MQTT messages, default is 2
		DisconnectTimeout uint // DisconnectTimeout wait the specified number of milliseconds for existing work to be completed before disconnecting.
		AvailabilityTopic topic.Topic
	}
)

func DefaultOptions() *ClientOptions {
	return &ClientOptions{
		Servers:           []*url.URL{},
		ClientID:          "",
		Username:          "",
		Password:          "",
		CleanSession:      true,
		QoS:               DefaultQoS,
		DisconnectTimeout: DefaultDisconnectTimeout, // Default disconnect timeout in milliseconds
	}
}

type (
	DialOptions func(*Conn)
	HandlerFunc func(ctx context.Context, conn model.StateUpdater, msg model.MQTTMessage) error
)

func WithOptions(opts *ClientOptions) DialOptions {
	return func(c *Conn) {
		c.opts = opts
	}
}

func WithHandler(handler model.Handler) DialOptions {
	return func(c *Conn) {
		c.handler = handler
	}
}

func WithOnConnectHandler(handler OnConnectHandler) DialOptions {
	return func(c *Conn) {
		c.onConnectHandler = func(_ mqtt.Client) {
			handler()
		}
	}
}

func WithOnConnectionLostHandler(handler OnConnectionLostHandler) DialOptions {
	return func(c *Conn) {
		c.onConnectionLostHandler = func(_ mqtt.Client, err error) {
			handler(err)
		}
	}
}

func WithDevice(dev *device.Device) DialOptions {
	return func(c *Conn) {
		c.dev = dev
	}
}

func WithOrigin(origin *origin.Origin) DialOptions {
	return func(c *Conn) {
		c.origin = origin
	}
}

func WithComponents(components ...*component.Component) DialOptions {
	return func(c *Conn) {
		c.components = append(c.components, components...)
	}
}

// func (c *ClientOptions) SetDiscoveryPrefix(prefix string) *ClientOptions {
// 	c.DiscoveryPrefix = prefix
// 	return c
// }

// func (c *ClientOptions) GetDeviceConfig() *deviceDiscoveryOptions {
// 	return &deviceDiscoveryOptions{c: c}
// }

func (c *ClientOptions) SetServers(server ...*url.URL) *ClientOptions {
	c.Servers = server
	return c
}

func (c *ClientOptions) AddServers(server ...*url.URL) *ClientOptions {
	c.Servers = append(c.Servers, server...)
	return c
}

func (c *ClientOptions) SetClientID(clientID string) *ClientOptions {
	c.ClientID = clientID
	return c
}

func (c *ClientOptions) SetUsername(username string) *ClientOptions {
	c.Username = username
	return c
}

func (c *ClientOptions) SetPassword(password string) *ClientOptions {
	c.Password = password
	return c
}

func (c *ClientOptions) SetCleanSession(cleanSession bool) *ClientOptions {
	c.CleanSession = cleanSession
	return c
}

func (c *ClientOptions) SetDefaultQoS(qos byte) *ClientOptions {
	c.QoS = qos
	return c
}

func (c *ClientOptions) SetAvailabilityTopic(topic topic.Topic) *ClientOptions {
	c.AvailabilityTopic = topic
	return c
}

// type deviceDiscoveryOptions struct {
// 	c *ClientOptions
// }

// func (d *deviceDiscoveryOptions) DiscoveryPrefix() string {
// 	return d.c.DiscoveryPrefix
// }
