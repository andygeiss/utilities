package messaging

import (
	"context"
	"sync"
)

// Bus ...
type Bus struct {
	subscribers []Actor
}

// Publish ...
func (a *Bus) Publish(ctx context.Context) {
	wg := sync.WaitGroup{}
	wg.Add(len(a.subscribers))
	for _, actor := range a.subscribers {
		go func(actor Actor) {
			actor.Receive(ctx)
			wg.Done()
		}(actor)
	}
	wg.Wait()
}

// Subscribe ...
func (a *Bus) Subscribe(actor Actor) {
	a.subscribers = append(a.subscribers, actor)
}

// NewBus ...
func NewBus() *Bus {
	return &Bus{
		subscribers: make([]Actor, 0),
	}
}
