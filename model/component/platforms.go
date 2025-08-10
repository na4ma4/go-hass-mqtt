package component

import "github.com/na4ma4/go-hass-mqtt/ptrval"

func AsSensor(c *Component) {
	c.Platform = ptrval.String("sensor")
}

func AsSwitch(c *Component) {
	c.Platform = ptrval.String("switch")
}

func AsText(c *Component) {
	c.Platform = ptrval.String("text")
}

func AsSelect(c *Component) {
	c.Platform = ptrval.String("select")
}

func AsNumber(c *Component) {
	c.Platform = ptrval.String("number")
}
