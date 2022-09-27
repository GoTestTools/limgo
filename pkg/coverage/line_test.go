package coverage_test

import (
	"testing"

	"github.com/engelmi/limgo/pkg/coverage"
)

func TestParsedLineKey(t *testing.T) {
	line := coverage.ParsedLine{"github.com/engelmi/limgo", "line.go", "11", "5", "11", "58", "1", "1"}
	expectedKey := "github.com/engelmi/limgo/line.go"
	if line.Key() != expectedKey {
		t.Fatalf("Expected line key to be '%s', but got '%s'", expectedKey, line.Key())
	}
}
