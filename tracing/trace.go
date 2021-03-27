package tracing

import (
	"context"
	"fmt"
)

type contextKey string

const (
	contextKeyTracing = contextKey("trace")
)

// Trace ...
type Trace struct {
	spans []*Span
	title string
}

// Add ...
func (a *Trace) Add(span *Span) *Trace {
	a.spans = append(a.spans, span)
	return a
}

// ToContext ...
func (a *Trace) ToContext(parent context.Context) context.Context {
	return context.WithValue(parent, contextKeyTracing, a)
}

// ToPlantUML ...
func (a *Trace) ToPlantUML() string {
	out := "@startuml " + a.title + "\n"
	for _, span := range a.spans {
		out += fmt.Sprintf("%s -> %s: %s (%v)\n", span.Source, span.Target, span.Label, span.Duration)
	}
	out += "@enduml\n"
	return out
}

// NewTrace ...
func NewTrace(title string) *Trace {
	return &Trace{
		spans: make([]*Span, 0),
		title: title,
	}
}

// FromContext ...
func FromContext(ctx context.Context) *Trace {
	switch trace := ctx.Value(contextKeyTracing).(type) {
	case *Trace:
		return trace
	}
	return NewTrace("trace")
}
