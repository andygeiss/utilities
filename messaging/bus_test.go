package messaging_test

import (
	"context"
	"testing"

	"github.com/andygeiss/utilities/messaging"
	assert "github.com/andygeiss/utilities/testing"
)

type ActorMessageDataStub1 struct{ Foo string }
type ActorMessageDataStub2 struct{ Bar string }

type ActorStub struct {
	Bus   *messaging.Bus
	State int
}

func (a *ActorStub) ID() string { return "ActorStub" }

func (a *ActorStub) Receive(ctx context.Context) {
	msg := messaging.FromContext(ctx)
	if msg == nil || msg.Data == nil {
		return
	}
	switch inbound := msg.Data.(type) {
	case *ActorMessageDataStub1:
		a.Bus.Publish(messaging.NewMessage(a.ID(), &ActorMessageDataStub2{Bar: inbound.Foo}).ToContext(ctx))
	case *ActorMessageDataStub2:
		a.State = 42
	}
}

func TestBus(t *testing.T) {
	bus := messaging.NewBus()
	actor := &ActorStub{Bus: bus, State: 0}
	bus.Subscribe(actor)
	bus.Publish(messaging.NewMessage(actor.ID(), &ActorMessageDataStub1{Foo: "foo"}).ToContext(context.Background()))
	assert.That("actor state should be 42", t, actor.State, 42)
}
