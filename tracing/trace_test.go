package tracing_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	assert "github.com/andygeiss/utilities/testing"
	"github.com/andygeiss/utilities/tracing"
)

func TestTraceAddAndToPlantUML(t *testing.T) {
	trace := tracing.NewTrace("title")
	trace.Add(tracing.NewSpan("source", "target", "label", time.Second))
	out := trace.ToPlantUML()
	wanted := `@startuml title
"source" -> "target": label (1s)
@enduml`
	assert.That("trace should not be nil", t, trace != nil, true)
	assert.That("out should be like wanted", t, out, wanted)
}

func TestToAndFromContext(t *testing.T) {
	trace := tracing.NewTrace("title")
	trace.Add(tracing.NewSpan("source", "target", "label", time.Second))
	ctx := trace.ToContext(context.Background())
	trace2 := tracing.FromContext(ctx)
	out := trace2.ToPlantUML()
	wanted := `@startuml title
"source" -> "target": label (1s)
@enduml`
	assert.That("trace should not be nil", t, trace != nil, true)
	assert.That("trace context should not be nil", t, ctx != nil, true)
	assert.That("trace2 should not be nil", t, trace2 != nil, true)
	assert.That("out should be like wanted", t, out, wanted)
}

func TestToAndFromContextTwice(t *testing.T) {
	trace := tracing.NewTrace("title")
	trace.Add(tracing.NewSpan("source", "target", "label", time.Second))
	ctx := trace.ToContext(context.Background())
	trace2 := tracing.FromContext(ctx)
	trace2.Add(tracing.NewSpan("source2", "target2", "label2", time.Second))
	out := trace2.ToPlantUML()
	wanted := `@startuml title
"source" -> "target": label (1s)
"source2" -> "target2": label2 (1s)
@enduml`
	assert.That("trace should not be nil", t, trace != nil, true)
	assert.That("trace context should not be nil", t, ctx != nil, true)
	assert.That("trace2 should not be nil", t, trace2 != nil, true)
	assert.That("out should be like wanted", t, out, wanted)
}

func TestFromTraceShouldHandleContextWithoutTrace(t *testing.T) {
	trace := tracing.FromContext(context.Background())
	assert.That("trace should not be nil", t, trace != nil, true)
}

func TestToFileShouldCreateFileStructure(t *testing.T) {
	ts := time.Now()
	path := filepath.Join("testdata")
	trace := tracing.FromContext(context.Background())
	fullPath := fmt.Sprintf("%s/%04d/%02d/%02d/%s.plantuml", path, ts.Year(), ts.Month(), ts.Day(), trace.Title)
	trace.ToFile(fullPath)
	stat, err := os.Stat(fullPath)
	assert.That("err should be nil", t, err, nil)
	assert.That("stat should not be nil", t, stat != nil, true)
	assert.That("testdata/YEAR/MONTH/DAY/TITLE.plantuml should be created", t, stat.IsDir(), true)
	os.RemoveAll("testdata")
}

func TestRegisterShouldCreateOneClient(t *testing.T) {
	trace := tracing.FromContext(context.Background())
	trace.Register("FooClient")
	out := trace.ToPlantUML()
	expected := `@startuml trace
entity "FooClient" #82b366
@enduml`
	assert.That("UML should be correct", t, out, expected)
}

func TestRegisterShouldCreateOneManager(t *testing.T) {
	trace := tracing.FromContext(context.Background())
	trace.Register("FooManager")
	out := trace.ToPlantUML()
	expected := `@startuml trace
entity "FooManager" #d6b656
@enduml`
	assert.That("UML should be correct", t, out, expected)
}

func TestRegisterShouldCreateOneEngine(t *testing.T) {
	trace := tracing.FromContext(context.Background())
	trace.Register("FooEngine")
	out := trace.ToPlantUML()
	expected := `@startuml trace
entity "FooEngine" #d79b00
@enduml`
	assert.That("UML should be correct", t, out, expected)
}

func TestRegisterShouldCreateOneResourceAccess(t *testing.T) {
	trace := tracing.FromContext(context.Background())
	trace.Register("FooResourceAccess")
	out := trace.ToPlantUML()
	expected := `@startuml trace
entity "FooResourceAccess" #6c8ebf
@enduml`
	assert.That("UML should be correct", t, out, expected)
}

func TestRegisterShouldCreateOneUnknownActor(t *testing.T) {
	trace := tracing.FromContext(context.Background())
	trace.Register("FooUnknown")
	out := trace.ToPlantUML()
	expected := `@startuml trace
entity "FooUnknown" #999999
@enduml`
	assert.That("UML should be correct", t, out, expected)
}
