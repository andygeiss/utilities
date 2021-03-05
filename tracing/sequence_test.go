package tracing_test

import (
	"testing"

	assert "github.com/andygeiss/utilities/testing"
	"github.com/andygeiss/utilities/tracing"
)

func TestSequenceTracerToPlantUMLWithoutATrace(t *testing.T) {
	wanted := `@startuml
@enduml`
	out := tracing.NewSequenceTracer().ToPlantUML()
	assert.That("trace should be empty", t, out, wanted)
}

func TestSequenceTracerToPlantUMLWithOneTraceSync(t *testing.T) {
	wanted := `@startuml
foo -> bar: message
@enduml`
	tracer := tracing.NewSequenceTracer()
	tracer.TraceSync("foo", "bar", "message")
	out := tracer.ToPlantUML()
	assert.That("trace should be empty", t, out, wanted)
}

func TestSequenceTracerToPlantUMLWithOneTraceAsync(t *testing.T) {
	wanted := `@startuml
foo ->> bar: message
@enduml`
	tracer := tracing.NewSequenceTracer()
	tracer.TraceAsync("foo", "bar", "message")
	out := tracer.ToPlantUML()
	assert.That("trace should be empty", t, out, wanted)
}
