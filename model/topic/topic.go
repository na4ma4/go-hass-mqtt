package topic

import "strings"

// Topic represents a topic in the MQTT broker.
// It provides methods to manipulate and format topics.
// This is used for device discovery and communication in the MQTT protocol.
type Topic string

const (
	// Root is the root topic for the MQTT broker.
	Root Topic = "homeassistant"
	// DeviceDiscovery is the topic for device discovery messages.
	DeviceDiscovery Topic = "homeassistant/device"
)

// New creates a new Topic from a string.
// It trims any leading or trailing whitespaces and slashes.
func New(topic string) Topic {
	topic = strings.TrimSpace(topic)
	topic = strings.Trim(topic, "/") // Remove leading and trailing slashes

	return Topic(topic)
}

// String returns the string representation of the Topic.
// It is used to convert the Topic to a string for MQTT operations.
func (t Topic) String() string {
	return string(t)
}

func (t Topic) Ptr() *Topic {
	return &t
}

// Wildcard returns a wildcard topic for matching all subtopics.
func (t Topic) Wildcard() Topic {
	if t == "" {
		return ""
	}
	return t.Join("#")
}

// Join concatenates the Topic with additional parts.
// It ensures that there are no trailing slashes and joins the parts with a slash.
func (t Topic) Join(parts ...string) Topic {
	if len(parts) == 0 {
		return t
	}

	base, _ := strings.CutSuffix(t.String(), "/") // Ensure no trailing slash

	v, _ := strings.CutSuffix(strings.Join(append([]string{base}, parts...), "/"), "/")

	return Topic(v)
}
