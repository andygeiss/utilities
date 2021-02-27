package messaging

// Bus defines the Publish/Subscribe pattern.
type Bus interface {
	Publish(message interface{})
	Subscribe(actor Actor)
}

type defaultBus struct {
	subscribers []Actor
}

// DefaultBus can be used to connect all the actors of the system.
var DefaultBus = NewBus()

// NewBus creates a new bus, which can be used to isolate different types of a bus.
func NewBus() Bus {
	return &defaultBus{
		subscribers: make([]Actor, 0),
	}
}

// Publish simply send the message to all the actors.
// Each actor is responsible for choosing and handling the relevant messages by itself.
func (a *defaultBus) Publish(message interface{}) {
	for _, actor := range a.subscribers {
		actor.Receive(message)
	}
}

// Subscribe simple registers an actor to the bus.
func (a *defaultBus) Subscribe(actor Actor) {
	a.subscribers = append(a.subscribers, actor)
}
