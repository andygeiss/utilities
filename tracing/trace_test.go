package tracing_test

import (
	"context"
	"testing"
	"time"

	assert "github.com/andygeiss/utilities/testing"
	"github.com/andygeiss/utilities/tracing"
)

func TestTraceAddAndToPlantUML(t *testing.T) {
	trace := tracing.NewTrace()
	trace.Add(tracing.NewSpan("source", "target", "label", time.Second))
	out := trace.ToPlantUML()
	wanted := `@startuml
source -> target: label
@enduml
`
	assert.That("trace should not be nil", t, trace != nil, true)
	assert.That("out should be like wanted", t, out, wanted)
}

func TestToAndFromContext(t *testing.T) {
	trace := tracing.NewTrace()
	trace.Add(tracing.NewSpan("source", "target", "label", time.Second))
	ctx := trace.ToContext(context.Background())
	trace2 := tracing.FromContext(ctx)
	out := trace2.ToPlantUML()
	wanted := `@startuml
source -> target: label
@enduml
`
	assert.That("trace should not be nil", t, trace != nil, true)
	assert.That("trace context should not be nil", t, ctx != nil, true)
	assert.That("trace2 should not be nil", t, trace2 != nil, true)
	assert.That("out should be like wanted", t, out, wanted)
}
