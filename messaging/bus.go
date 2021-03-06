package messaging

import (
	"context"
	"sync"
)

// Actor ...
type Actor func(ctx context.Context, in interface{})

// Message ...
type Message interface{}

// Bus ...
type Bus interface {
	Subscribe(topic string, actor Actor)
	Publish(ctx context.Context, topic string, msg Message)
}

type defaultBus struct {
	topics map[string][]Actor
}

// Subscribe ...
func (a *defaultBus) Subscribe(topic string, actor Actor) {
	a.topics[topic] = append(a.topics[topic], actor)
}

// Publish ...
func (a *defaultBus) Publish(ctx context.Context, topic string, msg Message) {
	if consumers, exists := a.topics[topic]; exists {
		wg := sync.WaitGroup{}
		wg.Add(len(consumers))
		for _, actor := range consumers {
			go func(actor Actor) {
				actor(ctx, msg)
				wg.Done()
			}(actor)
		}
		wg.Wait()
	}
}

// NewDefaultBus ...
func NewDefaultBus() Bus {
	return &defaultBus{
		topics: make(map[string][]Actor),
	}
}
