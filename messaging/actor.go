package messaging

import "context"

// Actor ...
type Actor interface {
	Receive(ctx context.Context)
}
