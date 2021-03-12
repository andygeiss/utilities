package messaging

import "context"

// Actor ...
type Actor interface {
	ID() string
	Receive(ctx context.Context)
}
