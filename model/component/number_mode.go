package component

import "github.com/na4ma4/go-hass-mqtt/ptrval"

func NumberModeAuto(c *Component) {
	c.Mode = ptrval.String("auto")
}

func NumberModeBox(c *Component) {
	c.Mode = ptrval.String("box")
}

func NumberModeSlider(c *Component) {
	c.Mode = ptrval.String("slider")
}
