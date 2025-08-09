package mqttconn

const (
	DefaultDisconnectTimeout uint = 500 // DefaultDisconnectTimeout is the default timeout in milliseconds for disconnecting the MQTT client.
	DefaultQoS               byte = 2   // DefaultQoS is the default Quality of Service level for MQTT messages.
)

var bsOnline = []byte("online")
