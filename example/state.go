//nolint:depguard // example files
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/na4ma4/go-hass-mqtt/model"
	"github.com/na4ma4/go-hass-mqtt/model/topic"
	"github.com/na4ma4/go-hass-mqtt/ptrval"
)

type exampleDevice struct {
	cmdTopic    topic.Topic `json:"-"`
	DisplayText *string     `json:"display_text,omitempty"`
	Time24Hour  *string     `json:"time_24_hour,omitempty"`
}

func NewExampleDevice(cmdTopic topic.Topic) *exampleDevice {
	return &exampleDevice{
		cmdTopic:    cmdTopic,
		DisplayText: ptrval.String("Hello, World! " + time.Now().Format(time.Kitchen)),
		Time24Hour:  ptrval.String("ON"),
	}
}

func (d *exampleDevice) HomeAssistantMessage(_ context.Context, _ model.StateUpdater, msg model.MQTTMessage) error {
	// Handle Home Assistant message
	log.Printf("Received Home Assistant message on topic %s: %s", msg.Topic(), msg.Payload())
	return nil
}

func (d *exampleDevice) HandleDeviceMessage(ctx context.Context, conn model.StateUpdater, msg model.MQTTMessage) error {
	// Handle Home Assistant message
	log.Printf("Received Device message on topic %s: %s", msg.Topic(), msg.Payload())

	if msg.Topic() != d.cmdTopic.String() {
		log.Printf("Ignoring message on topic %s, expected %s", msg.Topic(), d.cmdTopic.String())
		return nil
	}

	dm := &exampleDevice{}
	if err := json.Unmarshal(msg.Payload(), dm); err != nil {
		log.Printf("Error unmarshalling device message: %v", err)
		return nil
	}
	updateState := false
	// Update the device state based on the message
	if dm.DisplayText != nil {
		log.Printf("Updating display text to: %s", *dm.DisplayText)
		d.DisplayText = dm.DisplayText
		updateState = true
	}
	if dm.Time24Hour != nil {
		log.Printf("Updating time 24 hour to: %s", *dm.Time24Hour)
		d.Time24Hour = dm.Time24Hour
		updateState = true
	}

	if updateState {
		if err := conn.UpdateState(ctx); err != nil {
			log.Printf("Error updating state: %v", err)
		}
	}
	return nil
}

func (d *exampleDevice) State(_ context.Context) (*bytes.Buffer, error) {
	// Return the current state of the device
	buf := bytes.NewBuffer(nil)
	err := json.NewEncoder(buf).Encode(d)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
