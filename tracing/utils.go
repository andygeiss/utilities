package tracing

import (
	"context"
	"encoding/hex"
	"time"

	"github.com/andygeiss/utilities/security"
)

// NewRequestContextWithID ...
func NewRequestContextWithID() (ctx context.Context, id string) {
	key := security.NewKey256()
	id = hex.EncodeToString(key[:])
	return NewTrace(id).ToContext(context.Background()), id
}

// Call ...
func Call(ctx context.Context, source, target, action string, fn func()) context.Context {
	t0 := time.Now()
	// Local call should not be displayed with a request/response.
	if source == target {
		return FromContext(ctx).Add(NewSpan(source, target, action, time.Since(t0))).ToContext(ctx)
	}
	// Handle remote calls with a request/response.
	ctx = FromContext(ctx).Add(NewSpan(source, target, action, time.Since(t0))).ToContext(ctx)
	fn()
	return FromContext(ctx).Add(NewSpan(target, source, "", time.Since(t0))).ToContext(ctx)
}
