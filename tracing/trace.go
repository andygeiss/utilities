package tracing

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type contextKey string

const (
	contextKeyTracing = contextKey("trace")
)

// Trace ...
type Trace struct {
	spans []*Span
	Title string
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
	out := "@startuml " + a.Title + "\n"
	for _, span := range a.spans {
		out += fmt.Sprintf(`"%s" -> "%s": %s (%v)
`, span.Source, span.Target, span.Label, span.Duration)
	}
	out += "@enduml"
	return out
}

// ToFile ...
func (a *Trace) ToFile(path string) {
	ts := time.Now()
	fullPath := filepath.Join(path, strconv.Itoa(ts.Year()))
	fullPath = filepath.Join(fullPath, strconv.Itoa(int(ts.Month())))
	fullPath = filepath.Join(fullPath, strconv.Itoa(int(ts.Day())))
	os.MkdirAll(fullPath, 0755)
	ioutil.WriteFile(filepath.Join(fullPath, a.Title+".plantuml"), []byte(a.ToPlantUML()), 0644)
}

// NewTrace ...
func NewTrace(title string) *Trace {
	return &Trace{
		spans: make([]*Span, 0),
		Title: title,
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
