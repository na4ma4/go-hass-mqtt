package component

import "github.com/na4ma4/go-hass-mqtt/ptrval"

func ClassTemperature(c *Component) {
	c.DeviceClass = ptrval.String("temperature")
	c.UnitOfMeasurement = ptrval.String("°C")
}

func ClassHumidity(c *Component) {
	c.DeviceClass = ptrval.String("humidity")
	c.UnitOfMeasurement = ptrval.String("%")
}

func ClassPower(c *Component) {
	c.DeviceClass = ptrval.String("power")
	c.UnitOfMeasurement = ptrval.String("W")
}

func ClassVoltage(c *Component) {
	c.DeviceClass = ptrval.String("voltage")
	c.UnitOfMeasurement = ptrval.String("V")
}

func ClassCurrent(c *Component) {
	c.DeviceClass = ptrval.String("current")
	c.UnitOfMeasurement = ptrval.String("A")
}

func ClassEnergy(c *Component) {
	c.DeviceClass = ptrval.String("energy")
	c.UnitOfMeasurement = ptrval.String("kWh")
}

func ClassFrequency(c *Component) {
	c.DeviceClass = ptrval.String("frequency")
	c.UnitOfMeasurement = ptrval.String("Hz")
}

func ClassPowerFactor(c *Component) {
	c.DeviceClass = ptrval.String("power_factor")
	c.UnitOfMeasurement = ptrval.String("")
}

func ClassSignalStrength(c *Component) {
	c.DeviceClass = ptrval.String("signal_strength")
	c.UnitOfMeasurement = ptrval.String("dBm")
}
