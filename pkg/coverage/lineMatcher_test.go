package coverage_test

import (
	"testing"

	"github.com/GoTestTools/limgo/pkg/coverage"
)

func TestLineMatcher_Apply(t *testing.T) {

	testcases := []struct {
		name        string
		line        string
		expectError bool
	}{
		{
			name:        "fails on mode line",
			line:        "mode: atomic",
			expectError: true,
		},
		{
			name:        "successfully matches coverage line, example 1",
			line:        "github.com/GoTestTools/limgo/pkg/config/json.go:8.56,12.16 4 1",
			expectError: false,
		},
		{
			name:        "successfully matches coverage line, example 2",
			line:        "github.com/GoTestTools/limgo/pkg/config/json.go:12.16,14.3 1 0",
			expectError: false,
		},
		{
			name:        "successfully matches coverage line, example 3",
			line:        "github.com/GoTestTools/limgo/pkg/config/json.go:18.58,22.16 4 2",
			expectError: false,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			_, err := coverage.NewLineMatcher().Apply(testcase.line)
			if testcase.expectError && err == nil {
				t.Fatalf("Expected error, but got none")
			}
			if !testcase.expectError && err != nil {
				t.Fatalf("Expected no error, but got %v", err)
			}
		})
	}
}
