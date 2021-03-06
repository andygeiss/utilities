package messaging_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/andygeiss/utilities/messaging"
	assert "github.com/andygeiss/utilities/testing"
)

func TestBus(t *testing.T) {
	bus := messaging.NewDefaultBus()
	state := 0
	bus.Subscribe("foo", func(ctx context.Context, in interface{}) {
		state = 42
	})
	bus.Publish(context.Background(), "foo", "bar")
	assert.That("state should be changed to 42", t, state, 42)
}

type ActorStub struct {
	Bus   messaging.Bus
	State int
}

type BarMessage struct {
	Num int
}

type ErrorMessage struct {
	Error error
}

type FooMessage struct {
	Num int
}

type TimeoutMessage struct {
}

func (a *ActorStub) Foo(ctx context.Context, in interface{}) {
	done := make(chan struct{})
	// Handle Asynchronously ...
	var err error
	go func() {
		switch msg := in.(type) {
		case FooMessage:
			time.Sleep(time.Second)
			a.State = msg.Num * 2
		}
		done <- struct{}{}
	}()
	// Wait ...
	select {
	case <-ctx.Done(): // Context error
		a.Bus.Publish(ctx, "error", ErrorMessage{Error: ctx.Err()})
	case <-done:
		if err != nil { // Business error
			a.Bus.Publish(ctx, "error", ErrorMessage{Error: errors.New("error during foo")})
			return
		} // Success
		a.Bus.Publish(ctx, "bar", BarMessage{Num: a.State})
	}
}

func BenchmarkBus(b *testing.B) {
	bus := messaging.NewDefaultBus()
	for _, num := range []int{1, 2, 4, 8, 16, 32, 64, 128} {
		b.Run(fmt.Sprintf("%d", num), func(b *testing.B) {
			for i := 0; i < num; i++ {
				actor := &ActorStub{Bus: bus}
				bus.Subscribe("foo", actor.Foo)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				bus.Publish(context.Background(), "foo", FooMessage{Num: i})
			}
		})
	}
}

func TestBusStateChangeParallel(t *testing.T) {
	bus := messaging.NewDefaultBus()
	actors := make([]*ActorStub, 16)
	for i := 0; i < 16; i++ {
		actors[i] = &ActorStub{Bus: bus}
		bus.Subscribe("foo", actors[i].Foo)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second/2)
	defer cancel()
	bus.Publish(ctx, "foo", FooMessage{Num: 42})
	for i := 0; i < 16; i++ {
		assert.That(fmt.Sprintf("actor %d should changed its state to 84", i), t, actors[i].State, 84)
	}
}
