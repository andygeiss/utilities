package prototyping

import (
	"context"
	"sync"

	"github.com/andygeiss/utilities/logging"
)

// Bus defines the Publish/Subscribe pattern.
type Bus interface {
	Publish(message interface{})
	Subscribe(actor Actor)
}

type defaultBus struct {
	logger      logging.Logger
	subscribers []Actor
}

// Publish simply send the message to all the actors in parallel.
// Each actor is responsible for choosing and handling the relevant messages by itself.
func (a *defaultBus) Publish(message interface{}) {
	wg := sync.WaitGroup{}
	wg.Add(len(a.subscribers))
	for _, actor := range a.subscribers {
		go func(actor Actor) {
			actor.Send(context.Background(), message)
			wg.Done()
		}(actor)
	}
	wg.Wait()
}

// Subscribe simple registers an actor to the bus.
func (a *defaultBus) Subscribe(actor Actor) {
	a.subscribers = append(a.subscribers, actor)
}

// NewDefaultBus creates a new bus, which can be used to isolate different types of a bus.
func NewDefaultBus(logger logging.Logger) Bus {
	return &defaultBus{
		logger:      logger,
		subscribers: make([]Actor, 0),
	}
}
