package mqttconn

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/na4ma4/go-contextual"
	"github.com/na4ma4/go-hass-mqtt/model"
	"github.com/na4ma4/go-hass-mqtt/model/component"
	"github.com/na4ma4/go-hass-mqtt/model/device"
	"github.com/na4ma4/go-hass-mqtt/model/messages"
	"github.com/na4ma4/go-hass-mqtt/model/origin"
	"github.com/na4ma4/go-hass-mqtt/model/topic"
	"golang.org/x/sync/errgroup"
)

type Conn struct {
	dev        *device.Device
	origin     *origin.Origin
	components []*component.Component
	handler    model.Handler
	conn       mqtt.Client
	opts       *ClientOptions
}

var ErrInvalidServer = errors.New("invalid server configuration")

func Dial(opts ...DialOptions) (*Conn, error) {
	// Create a new MQTT client with the provided options

	conn := &Conn{
		opts: DefaultOptions(),
	}

	for _, opt := range opts {
		opt(conn)
	}

	if len(conn.opts.Servers) == 0 {
		return nil, ErrInvalidServer
	}

	mqOpts := mqtt.NewClientOptions()
	for _, server := range conn.opts.Servers {
		if server == nil {
			return nil, ErrInvalidServer
		}
		mqOpts.AddBroker(server.String())
	}
	mqOpts.SetClientID(conn.opts.ClientID)
	mqOpts.SetUsername(conn.opts.Username)
	mqOpts.SetPassword(conn.opts.Password)
	mqOpts.SetCleanSession(conn.opts.CleanSession)
	conn.conn = mqtt.NewClient(mqOpts)

	token := conn.conn.Connect()
	if token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	// Implement the dialing logic here
	return conn, nil
}

func (c *Conn) Close() error {
	// Implement the closing logic here
	c.conn.Disconnect(c.opts.DisconnectTimeout)
	return nil
}

// Listener starts listening for MQTT messages and handles them using the provided handler.
func (c *Conn) Listener(ctx context.Context) error {
	if c.conn == nil {
		return errors.New("connection is not established")
	}

	if c.handler == nil {
		return errors.New("handler is not set")
	}

	errg, ctx := errgroup.WithContext(ctx)
	conCtx := contextual.NewCancellable(ctx)
	errg.Go(func() error {
		// Subscribe to the homeassistant topic
		return c.subscribeWithHandler(
			conCtx,
			c.dev.TopicHomeAssistant().Wildcard(),
			c.handler.HomeAssistantMessage,
		)
	})
	errg.Go(func() error {
		// Subscribe to the device topic
		return c.subscribeWithHandler(
			conCtx,
			c.dev.TopicDeviceBase().Wildcard(),
			c.handler.HandleDeviceMessage,
		)
	})

	return errg.Wait()
}

func (c *Conn) subscribeWithHandler(ctx contextual.Context,
	topic topic.Topic,
	handler HandlerFunc,
) error {
	if topic == "" {
		return errors.New("invalid topic")
	}
	tokenChan := make(chan mqtt.Message, 1)

	go func() {
		for {
			select {
			case msg := <-tokenChan:
				if err := handler(ctx, c, msg); err != nil {
					ctx.CancelWithCause(err)
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	{
		token := c.conn.Subscribe(
			topic.String(),
			c.opts.QoS,
			func(_ mqtt.Client, msg mqtt.Message) {
				tokenChan <- msg
			},
		)
		if token.Wait() && token.Error() != nil {
			return token.Error()
		}
	}

	// Wait for the context to be done before unsubscribing
	<-ctx.Done()
	if token := c.conn.Unsubscribe(topic.String()); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return ctx.Err()
}

func (c *Conn) sendMessage(_ context.Context, topic string, payload *bytes.Buffer) error {
	if c.conn == nil {
		return errors.New("connection is not established")
	}

	if topic == "" {
		return errors.New("invalid topic")
	}

	if payload == nil {
		return errors.New("payload cannot be nil")
	}

	token := c.conn.Publish(topic, c.opts.QoS, false, *payload)
	if token.Wait() && token.Error() != nil {
		return fmt.Errorf("sendMessage(): token error: %w", token.Error())
	}

	return nil
}

func (c *Conn) Connect(ctx context.Context) error {
	if c.conn == nil {
		return errors.New("connection is not established")
	}

	topic := c.dev.TopicHomeAssistant().Join("config")
	if topic == "" {
		return errors.New("invalid device homeassistant topic")
	}

	msg := messages.NewDiscovery(
		c.dev, c.origin,
		messages.WithCommandTopic(c.dev.TopicDeviceBase().Join("cmd").String()),
		messages.WithStateTopic(c.dev.TopicDeviceBase().Join("state").String()),
		messages.WithComponents(c.components...),
		messages.WithQos(DefaultQoS),
	)

	payload, err := msg.BytesBuffer()
	if err != nil {
		return err
	}

	return c.sendMessage(ctx, topic.String(), payload)
}

func (c *Conn) UpdateState(ctx context.Context) error {
	if c.conn == nil {
		return errors.New("connection is not established")
	}

	topic := c.dev.TopicDeviceBase().Join("state")
	if topic == "" {
		return errors.New("invalid device state topic")
	}

	payload, err := c.handler.State(ctx)
	if err != nil {
		return err
	}

	return c.sendMessage(ctx, topic.String(), payload)
}
