package tracing

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type contextKey string

const (
	contextKeyTracing = contextKey("trace")
)

// Trace ...
type Trace struct {
	actors []string
	spans  []*Span
	Title  string
}

// Add ...
func (a *Trace) Add(span *Span) *Trace {
	a.spans = append(a.spans, span)
	return a
}

// Register ...
func (a *Trace) Register(actor string) {
	exists := false
	for _, current := range a.actors {
		if current == actor {
			exists = true
			break
		}
	}
	if !exists {
		a.actors = append(a.actors, actor)
	}
}

// Spans ...
func (a *Trace) Spans() []*Span {
	return a.spans
}

// ToContext ...
func (a *Trace) ToContext(parent context.Context) context.Context {
	return context.WithValue(parent, contextKeyTracing, a)
}

// ToPlantUML ...
func (a *Trace) ToPlantUML() string {
	out := "@startuml " + a.Title + "\n"
	for _, actor := range a.actors {
		out += fmt.Sprintf(`entity "%s" %s
`, actor, getActorColor(actor))
	}
	for _, span := range a.spans {
		out += fmt.Sprintf(`"%s" -> "%s": %s (%v)
`, span.Source, span.Target, span.Label, span.Duration)
	}
	out += "@enduml"
	return out
}

func getActorColor(actor string) string {
	if strings.Contains(actor, "Client") {
		return "#82b366"
	} else if strings.Contains(actor, "Manager") {
		return "#d6b656"
	} else if strings.Contains(actor, "Engine") {
		return "#d79b00"
	} else if strings.Contains(actor, "ResourceAccess") {
		return "#6c8ebf"
	}
	return "#999999"
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
