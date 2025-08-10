package device

import "github.com/na4ma4/go-hass-mqtt/model/topic"

type OptFunc func(*Device)

// func WithConnection(discoveryTopic topic.Topic) OptFunc {
// 	return func(d *Device) {
// 		d.config.DiscoveryPrefix = config.DiscoveryPrefix()
// 		d.config.Component = "device"
// 		d.config.ObjectID = d.ID.String()
// 	}
// }

// func WithNodeID(nodeID string) OptFunc {
// 	return func(d *Device) {
// 		d.config.NodeID = nodeID
// 	}
// }

func WithDeviceTopic(topic topic.Topic) OptFunc {
	return func(d *Device) {
		d.baseDeviceTopic = &topic
	}
}

func WithHomeAssistantTopic(topic topic.Topic) OptFunc {
	return func(d *Device) {
		d.hassTopic = &topic
	}
}

func WithAvailabilityTopic(topic topic.Topic) OptFunc {
	return func(d *Device) {
		d.AvailabilityTopic = &topic
	}
}

func WithAvailabilityTemplate(template string) OptFunc {
	return func(d *Device) {
		d.AvailabilityTemplate = template
	}
}

func WithDeviceName(name string) OptFunc {
	return func(d *Device) {
		d.Name = name
	}
}

func WithDeviceManufacturer(manufacturer string) OptFunc {
	return func(d *Device) {
		d.Manufacturer = manufacturer
	}
}

func WithDeviceModel(model string) OptFunc {
	return func(d *Device) {
		d.Model = model
	}
}

func WithDeviceSoftwareVersion(software string) OptFunc {
	return func(d *Device) {
		d.SoftwareVersion = software
	}
}

func WithDeviceHardwareVersion(hardware string) OptFunc {
	return func(d *Device) {
		d.HardwareVersion = hardware
	}
}

func WithDeviceSerial(serial string) OptFunc {
	return func(d *Device) {
		d.Serial = serial
	}
}
