package messaging_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/andygeiss/utilities/logging"
	"github.com/andygeiss/utilities/messaging"
	assert "github.com/andygeiss/utilities/testing"
)

type ActorStub struct {
	Bus          messaging.Bus
	StateChanged bool
}

func (a *ActorStub) Name() string {
	return "ActorStub"
}

func (a *ActorStub) Receive(message interface{}) {
	switch message.(type) {
	case MessageStub:
		a.StateChanged = true
	}
	time.Sleep(time.Millisecond)
}

func (a *ActorStub) Send(message interface{}) {
	a.Bus.Publish(message)
}

type MessageStub struct{}

func TestBusPublishAfterSubscribe(t *testing.T) {
	logger := logging.NewDefaultLogger()
	bus := messaging.NewDefaultBus(logger)
	actor := &ActorStub{Bus: bus}
	bus.Subscribe(actor)
	bus.Publish(MessageStub{})
	assert.That("state of actor should be changed", t, actor.StateChanged, true)
}

func TestBusPublishWithoutSubscribe(t *testing.T) {
	logger := logging.NewDefaultLogger()
	bus := messaging.NewDefaultBus(logger)
	actor := &ActorStub{Bus: bus}
	bus.Publish(MessageStub{})
	assert.That("state of actor should not be changed", t, actor.StateChanged, false)
}

func TestBusWithTwoActorsBusSend(t *testing.T) {
	logger := logging.NewDefaultLogger()
	bus := messaging.NewDefaultBus(logger)
	actor1 := &ActorStub{Bus: bus}
	actor2 := &ActorStub{Bus: bus}
	bus.Subscribe(actor1)
	bus.Subscribe(actor2)
	bus.Publish(MessageStub{})
	assert.That("state of actor1 should be changed", t, actor1.StateChanged, true)
	assert.That("state of actor2 should be changed", t, actor2.StateChanged, true)
}

func TestBusWithTwoActorsActorSend(t *testing.T) {
	logger := logging.NewDefaultLogger()
	bus := messaging.NewDefaultBus(logger)
	actor1 := &ActorStub{Bus: bus}
	actor2 := &ActorStub{Bus: bus}
	bus.Subscribe(actor1)
	bus.Subscribe(actor2)
	actor1.Send(MessageStub{})
	assert.That("state of actor1 should be changed", t, actor1.StateChanged, true)
	assert.That("state of actor2 should be changed", t, actor2.StateChanged, true)
}

func BenchmarkBus(b *testing.B) {
	logger := logging.NewDefaultLogger()
	bus := messaging.NewDefaultBus(logger)
	numActors := []int{1, 2, 4, 8, 16, 1024, 4096}
	for _, num := range numActors {
		b.Run(fmt.Sprintf("%d actors", num), func(b *testing.B) {
			actors := make([]messaging.Actor, num)
			// Setup ...
			for i := 0; i < num; i++ {
				actor := &ActorStub{Bus: bus}
				actors[i] = actor
				bus.Subscribe(actor)
			}
			// Run ...
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				// The following message will be send to all actors.
				bus.Publish(MessageStub{})
			}
		})
	}
}
