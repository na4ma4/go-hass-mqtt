package model

import (
	"bytes"
	"context"
)

type StateUpdater interface {
	// UpdateState publishes the current state of the device to the MQTT broker.
	// It returns an error if the state could not be published.
	UpdateState(ctx context.Context) error
}

type Handler interface {
	HomeAssistantMessage(ctx context.Context, conn StateUpdater, msg MQTTMessage) error
	HandleDeviceMessage(ctx context.Context, conn StateUpdater, msg MQTTMessage) error
	State(ctx context.Context) (*bytes.Buffer, error)
}

// MQTTMessage represents a message received from the MQTT broker.
// It provides methods to access the message's properties such as topic, payload, QoS,
// and whether it is a duplicate or retained message.
// This interface is used for handling messages in the MQTT connection.
type MQTTMessage interface {
	Duplicate() bool
	Qos() byte
	Retained() bool
	Topic() string
	MessageID() uint16
	Payload() []byte
	Ack()
}
