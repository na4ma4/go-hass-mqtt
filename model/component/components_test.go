package component_test

import (
	"testing"

	"github.com/na4ma4/go-hass-mqtt/model"
	"github.com/na4ma4/go-hass-mqtt/model/component"
	"github.com/na4ma4/go-hass-mqtt/model/topic"
)

func TestNewComponent_Defaults(t *testing.T) {
	id := model.BasicIdentifier("test-id")
	c := component.New(&id)
	if c.ID == nil || *c.ID != id {
		t.Errorf("expected ID to be set")
	}
	if c.Name != nil {
		t.Errorf("expected Name to be nil")
	}
}

func TestWithName(t *testing.T) {
	id := model.BasicIdentifier("id")
	name := "TestName"
	c := component.New(&id, component.WithName(name))
	if c.Name == nil || *c.Name != name {
		t.Errorf("component.WithName did not set Name correctly")
	}
}

func TestWithBaseTopic(t *testing.T) {
	id := model.BasicIdentifier("id")
	topicVal := topic.Topic("base/topic")
	c := component.New(&id, component.WithBaseTopic(topicVal))
	if c.BaseTopic == nil || *c.BaseTopic != topicVal {
		t.Errorf("component.WithBaseTopic did not set BaseTopic correctly")
	}
}

func TestWithCommandTemplate(t *testing.T) {
	id := model.BasicIdentifier("id")
	template := "cmd_template"
	c := component.New(&id, component.WithCommandTemplate(template))
	if c.CommandTemplate == nil || *c.CommandTemplate != template {
		t.Errorf("component.WithCommandTemplate did not set CommandTemplate correctly")
	}
}

func TestWithValueTemplate(t *testing.T) {
	id := model.BasicIdentifier("id")
	template := "val_template"
	c := component.New(&id, component.WithValueTemplate(template))
	if c.ValueTemplate == nil || *c.ValueTemplate != template {
		t.Errorf("component.WithValueTemplate did not set ValueTemplate correctly")
	}
}

func TestWithCommandTopic(t *testing.T) {
	id := model.BasicIdentifier("id")
	cmdTopic := topic.Topic("cmd/topic")
	c := component.New(&id, component.WithCommandTopic(cmdTopic))
	if c.CommandTopic == nil || *c.CommandTopic != cmdTopic {
		t.Errorf("component.WithCommandTopic did not set CommandTopic correctly")
	}
}

func TestWithStateTopic(t *testing.T) {
	id := model.BasicIdentifier("id")
	stateTopic := topic.Topic("state/topic")
	c := component.New(&id, component.WithStateTopic(stateTopic))
	if c.StateTopic == nil || *c.StateTopic != stateTopic {
		t.Errorf("component.WithStateTopic did not set StateTopic correctly")
	}
}

func TestMultipleOptions(t *testing.T) {
	id := model.BasicIdentifier("id")
	name := "multi"
	base := topic.Topic("base")
	cmd := topic.Topic("cmd")
	state := topic.Topic("state")
	valTpl := "val"
	cmdTpl := "cmd"

	c := component.New(&id,
		component.WithName(name),
		component.WithBaseTopic(base),
		component.WithCommandTopic(cmd),
		component.WithStateTopic(state),
		component.WithValueTemplate(valTpl),
		component.WithCommandTemplate(cmdTpl),
	)

	if c.Name == nil || *c.Name != name {
		t.Errorf("Name not set")
	}
	if c.BaseTopic == nil || *c.BaseTopic != base {
		t.Errorf("BaseTopic not set")
	}
	if c.CommandTopic == nil || *c.CommandTopic != cmd {
		t.Errorf("CommandTopic not set")
	}
	if c.StateTopic == nil || *c.StateTopic != state {
		t.Errorf("StateTopic not set")
	}
	if c.ValueTemplate == nil || *c.ValueTemplate != valTpl {
		t.Errorf("ValueTemplate not set")
	}
	if c.CommandTemplate == nil || *c.CommandTemplate != cmdTpl {
		t.Errorf("CommandTemplate not set")
	}
}
