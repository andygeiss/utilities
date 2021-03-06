package prototyping_test

import (
	"context"
	"testing"
	"time"

	"github.com/andygeiss/utilities/prototyping"
	assert "github.com/andygeiss/utilities/testing"
)

func TestContextWithTimeout(t *testing.T) {
	// Arrange ...
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	actor := prototyping.NewDefaultActor()
	// Act ...
	res, err := actor.Request(ctx, "foo")
	// Assert ...
	assert.That("error should be nil", t, err, nil)
	assert.That("response should not be nil", t, res == nil, false)
}
