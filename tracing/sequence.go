package tracing

import (
	"fmt"
	"sync"
)

// SequenceTracer ...
type SequenceTracer struct {
	Flow  map[string]string
	mutex sync.Mutex
}

// ToPlantUML ...
func (a *SequenceTracer) ToPlantUML() string {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	var out string
	out += "@startuml\n"
	for k, v := range a.Flow {
		out += fmt.Sprintf("%s: %s\n", k, v)
	}
	out += "@enduml"
	return out
}

// TraceAsync ...
func (a *SequenceTracer) TraceAsync(source, target, message string) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.Flow[fmt.Sprintf("%s ->> %s", source, target)] = message
}

// TraceSync ...
func (a *SequenceTracer) TraceSync(source, target, message string) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.Flow[fmt.Sprintf("%s -> %s", source, target)] = message
}

// NewSequenceTracer ...
func NewSequenceTracer() *SequenceTracer {
	return &SequenceTracer{
		Flow: make(map[string]string, 0),
	}
}
