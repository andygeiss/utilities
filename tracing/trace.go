package tracing

import (
	"context"
	"fmt"
	"time"
)

type contextKey string

const (
	contextKeyTracing = contextKey("trace")
)

// Span ...
type Span struct {
	Duration time.Duration `json:"duration"`
	Label    string        `json:"label"`
	Source   string        `json:"source"`
	Target   string        `json:"target"`
}

// NewSpan ...
func NewSpan(source, target, label string, duration time.Duration) *Span {
	return &Span{
		Duration: duration,
		Label:    label,
		Source:   source,
		Target:   target,
	}
}

// Trace ...
type Trace struct {
	spans []*Span
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
	out := "@startuml\n"
	for _, span := range a.spans {
		out += fmt.Sprintf("%s -> %s: %s (%v)\n", span.Source, span.Target, span.Label, span.Duration)
	}
	out += "@enduml\n"
	return out
}

// NewTrace ...
func NewTrace() *Trace {
	return &Trace{
		spans: make([]*Span, 0),
	}
}

// FromContext ...
func FromContext(ctx context.Context) *Trace {
	switch trace := ctx.Value(contextKeyTracing).(type) {
	case *Trace:
		return trace
	}
	return nil
}
