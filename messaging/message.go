package messaging

import (
	"context"
	"encoding/hex"

	"github.com/andygeiss/utilities/security"
)

type contextKey string

const (
	contextKeyMessage = contextKey("message")
)

// Message ...
type Message struct {
	ID     string      `json:"id"`
	Data   interface{} `json:"data"`
	Source string      `json:"source"`
}

// ToContext ...
func (a *Message) ToContext(parent context.Context) context.Context {
	return context.WithValue(parent, contextKeyMessage, a)
}

// NewMessage ...
func NewMessage(source string, data interface{}) *Message {
	id := security.NewKey256()
	return &Message{
		ID:     hex.EncodeToString(id[:]),
		Data:   data,
		Source: source,
	}
}

// FromContext ...
func FromContext(ctx context.Context) *Message {
	switch msg := ctx.Value(contextKeyMessage).(type) {
	case *Message:
		return msg
	}
	return nil
}
