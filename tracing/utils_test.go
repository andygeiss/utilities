package tracing_test

import (
	"context"
	"testing"

	assert "github.com/andygeiss/utilities/testing"
	"github.com/andygeiss/utilities/tracing"
)

func TestLocalCallShouldAddOneSpan(t *testing.T) {
	value := 0
	ctx := tracing.Call(context.Background(), "source", "source", "local action", func() { value = 1 })
	spans := tracing.FromContext(ctx).Spans()
	assert.That("one span should be added", t, len(spans), 1)
	assert.That("value should be changed to 1", t, value, 1)
}

func TestRemoteCallShouldAddTwoSpans(t *testing.T) {
	value := 0
	ctx := tracing.Call(context.Background(), "source", "target", "remote action", func() { value = 1 })
	spans := tracing.FromContext(ctx).Spans()
	assert.That("two spans should be added", t, len(spans), 2)
	assert.That("value should be changed to 1", t, value, 1)
}

func TestNewRequestContextShouldAddAnRequestID(t *testing.T) {
	ctx, id := tracing.NewRequestContextWithID()
	trace := tracing.FromContext(ctx)
	assert.That("id should not be empty string", t, id == "", false)
	assert.That("title should be the id", t, id == trace.Title, true)
}
