package messaging

// Actor defines a simple interface to ensure that each actor is able to choose and handle relevant messages by itself.
type Actor interface {
	Receive(message interface{})
}
