package messaging_test

import (
	"testing"

	"github.com/andygeiss/utilities/messaging"
	assert "github.com/andygeiss/utilities/testing"
)

type ActorStub struct {
	Bus          messaging.Bus
	StateChanged bool
}

func (a *ActorStub) Receive(message interface{}) {
	switch message.(type) {
	case MessageStub:
		a.StateChanged = true
	}
}

type MessageStub struct{}

func TestBusPublishAfterSubscribe(t *testing.T) {
	bus := messaging.NewBus()
	actor := &ActorStub{Bus: bus}
	bus.Subscribe(actor)
	bus.Publish(MessageStub{})
	assert.That("state of actor should be changed", t, actor.StateChanged, true)
}

func TestBusPublishWithoutSubscribe(t *testing.T) {
	bus := messaging.NewBus()
	actor := &ActorStub{Bus: bus}
	bus.Publish(MessageStub{})
	assert.That("state of actor should not be changed", t, actor.StateChanged, false)
}

func TestTwoActors(t *testing.T) {
	bus := messaging.NewBus()
	actor1 := &ActorStub{Bus: bus}
	actor2 := &ActorStub{Bus: bus}
	bus.Subscribe(actor1)
	bus.Subscribe(actor2)
	bus.Publish(MessageStub{})
	assert.That("state of actor1 should be changed", t, actor1.StateChanged, true)
	assert.That("state of actor2 should be changed", t, actor2.StateChanged, true)
}
