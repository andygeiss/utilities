package messaging_test

import (
	"context"
	"testing"

	"github.com/andygeiss/utilities/messaging"
	assert "github.com/andygeiss/utilities/testing"
)

type MessageStub struct{ Text string }

func TestNewMessage(t *testing.T) {
	msg := messaging.NewMessage("foo", &MessageStub{Text: "foo"})
	assert.That("message should not be nil", t, msg != nil, true)
	assert.That("message data should not be nil", t, msg.Data != nil, true)
	assert.That("message data text should be foo", t, msg.Data.(*MessageStub).Text, "foo")
}

func TestToAndFromContext(t *testing.T) {
	msg := messaging.NewMessage("foo", &MessageStub{Text: "foo"})
	ctx := msg.ToContext(context.Background())
	msg2 := messaging.FromContext(ctx)
	assert.That("message2 should not be nil", t, msg2 != nil, true)
	assert.That("message2 data should not be nil", t, msg2.Data != nil, true)
	assert.That("message2 data text should be foo", t, msg2.Data.(*MessageStub).Text, "foo")
}
