package prototyping_test

import (
	"testing"

	"github.com/andygeiss/utilities/prototyping"
	assert "github.com/andygeiss/utilities/testing"
)

func TestContextWithValue(t *testing.T) {
	// Create a context with a message
	ctx := prototyping.NewMessage("foo").ToContext()
	// Extract the message from the context
	msg, err := prototyping.FromContext(ctx)
	assert.That("error should be nil", t, err, nil)
	assert.That("message should not be nil", t, msg == nil, false)
	assert.That("message data should be foo", t, msg.Data, "foo")
}
