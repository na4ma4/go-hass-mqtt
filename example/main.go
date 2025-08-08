package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/na4ma4/go-contextual"
	"github.com/na4ma4/go-hass-mqtt/model"
	"github.com/na4ma4/go-hass-mqtt/model/component"
	"github.com/na4ma4/go-hass-mqtt/model/device"
	"github.com/na4ma4/go-hass-mqtt/model/origin"
	"github.com/na4ma4/go-hass-mqtt/model/topic"
	"github.com/na4ma4/go-hass-mqtt/mqttconn"
	"github.com/na4ma4/go-hass-mqtt/ptrval"
)

//nolint:funlen // main function to run the example
func main() {
	ctx := contextual.NewCancellable(
		context.Background(),
		contextual.WithSignalCancelOption(
			os.Interrupt,
		),
	)
	defer ctx.Cancel()

	org := origin.New(
		"example-project",
		"1.0.0",
		"https://github.com/na4ma4/go-hass-mqtt",
	)

	deviceTopic := topic.New("example-project").Join(
		"example-device",
	)
	homeAssistantBaseTopic := topic.New("homeassistant").Join(
		"device",
		"example-project",
		"example-device",
	)

	dev := device.New("example-device",
		device.WithDeviceTopic(deviceTopic),
		device.WithHomeAssistantTopic(homeAssistantBaseTopic),
		device.WithDeviceName("Test Device"),
		device.WithDeviceManufacturer("Test Manufacturer"),
		device.WithDeviceModel("Test Model"),
		device.WithDeviceSoftwareVersion("1.0.0"),
		device.WithDeviceHardwareVersion("1.0rev1"),
		device.WithDeviceSerial("Test Serial"),
	)

	cmps := []*component.Component{
		component.New(model.BasicIdentifier(dev.ID.String()+"_display").Ptr(),
			component.AsText,
			component.WithValueTemplate("{{ value_json.display_text }}"),
			component.WithName("Display Text"),
			component.WithCommandTemplate(`{"display_text": "{{ value }}"}`),
		),
		component.New(model.BasicIdentifier(dev.ID.String()+"_24hr").Ptr(),
			component.AsSwitch,
			component.WithValueTemplate("{{ value_json.time_24_hour }}"),
			component.WithCommandTemplate(`{"time_24_hour": "{{ value }}"}`),
			component.WithName("24 Hour Mode"),
		),
	}

	handler := NewExampleDevice(
		deviceTopic.Join("cmd"),
	)

	connOpt := mqttconn.DefaultOptions().
		SetServers(ptrval.MustURL("tcp://homeassistant.local:1883")).
		SetClientID("test_client_id_example").
		SetUsername(os.Getenv("MQTT_USERNAME")).
		SetPassword(os.Getenv("MQTT_PASSWORD"))

	var conn *mqttconn.Conn
	{
		var err error
		conn, err = mqttconn.Dial(
			mqttconn.WithOptions(connOpt),
			mqttconn.WithHandler(handler),
			mqttconn.WithDevice(dev),
			mqttconn.WithOrigin(org),
			mqttconn.WithComponents(cmps...),
		)
		if err != nil {
			panic(err)
		}
		defer conn.Close()
	}

	if err := conn.Connect(ctx); err != nil {
		log.Printf("Failed to connect: %v", err)
		return
	}

	ctx.Go(func() error {
		action := func() error {
			if err := conn.UpdateState(ctx); err != nil {
				log.Printf("Failed to publish state message: %v", err)
				return err
			}
			return nil
		}

		if err := action(); err != nil {
			log.Printf("Failed to publish state message: %v", err)
			return err
		}

		ticker := time.NewTicker(time.Second * 10) //nolint:mnd // example purpose
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if err := action(); err != nil {
					log.Printf("Failed to publish state message: %v", err)
					return err
				}
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	})

	ctx.Go(func() error {
		return conn.Listener(ctx)
	})

	if err := ctx.Wait(); err != nil {
		log.Printf("routine returned error: %s", err)
		return
	}
}
