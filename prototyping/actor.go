package prototyping

import (
	"context"
	"errors"
)

// Actor ...
type Actor interface {
	Receive(req interface{}) (res interface{}, err error)
	Request(ctx context.Context, req interface{}) (res interface{}, err error)
	Send(ctx context.Context, req interface{})
}

type defaultActor struct{}

// Request ...
func (a *defaultActor) Request(ctx context.Context, req interface{}) (res interface{}, err error) {
	done := make(chan struct{})
	// Handle Request/Response ...
	go func() {
		res, err = a.Receive(req)
		done <- struct{}{}
	}()
	// Wait ...
	select {
	case <-ctx.Done(): // Timeout ...
		return nil, ctx.Err()
	case <-done: // No Timeout ...
		if err != nil { // Error handling ...
			return nil, err
		}
		return res, nil // Success ...
	}
}

// Send ...
func (a *defaultActor) Send(ctx context.Context, req interface{}) {
	// Fire and forget ...
	go func() {
		a.Request(ctx, req)
	}()
}

// Receive ...
func (a *defaultActor) Receive(req interface{}) (res interface{}, err error) {
	return nil, errors.New("not implemented")
}

// NewDefaultActor ...
func NewDefaultActor() Actor {
	return &defaultActor{}
}
