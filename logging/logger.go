package logging

import (
	"log"
	"os"
)

// Logger ...
type Logger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
}

// NewDefaultLogger ...
func NewDefaultLogger() Logger {
	return log.New(os.Stdout, "", log.Ldate|log.Ltime)

}
