package prototyping

import (
	"context"
	"errors"
)

type contextKey string

func (a contextKey) String() string {
	return "context key " + string(a)
}

var (
	contextKeyMessage = contextKey("message")
)

// Message ...
type Message struct {
	Data interface{}
}

// ToContext ...
func (a *Message) ToContext() context.Context {
	return context.WithValue(context.Background(), contextKeyMessage, a)
}

// ToContextWithParent ...
func (a *Message) ToContextWithParent(ctx context.Context) context.Context {
	return context.WithValue(ctx, contextKeyMessage, a)
}

// FromContext ...
func FromContext(ctx context.Context) (msg *Message, err error) {
	switch val := ctx.Value(contextKeyMessage).(type) {
	case *Message:
		return val, nil
	default:
		return nil, errors.New("message not found")
	}
}

// NewMessage ...
func NewMessage(data interface{}) *Message {
	return &Message{Data: data}
}
