package mqttconn

const (
	DefaultDisconnectTimeout uint = 1000 // DefaultDisconnectTimeout is the default timeout in milliseconds for disconnecting the MQTT client.
	DefaultQoS               byte = 2    // DefaultQoS is the default Quality of Service level for MQTT messages.
)

const (
	LWTHello string = "OK"
	LWTBye   string = "GONE"
)

//nolint:gochecknoglobals // defined once constant.
var bsOnline = []byte("online")
