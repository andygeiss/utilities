package tracing

import "time"

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
