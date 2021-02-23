package assert

import (
	"fmt"
	"testing"
)

// That describes a specific test-case with by a description [desc].
func That(desc string, t *testing.T, got interface{}, exp interface{}) {
	t.Run(desc, func(t *testing.T) {
		if fmt.Sprintf("%v", exp) != fmt.Sprintf("%v", got) {
			t.Errorf("expected [%v] but got [%v]", exp, got)
		}
	})
}
