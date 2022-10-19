package domain_test

import (
	"os"
	"path"
	"testing"

	"github.com/GoTestTools/limgo/pkg/domain"
)

func TestParse(t *testing.T) {

	testcases := []struct {
		name     string
		file     string
		expected int
	}{
		{
			name:     "Parsing completely empty coverage file",
			file:     "completely_empty.cov",
			expected: 0,
		},
		{
			name:     "Parsing coverage file with no run tests",
			file:     "empty.cov",
			expected: 0,
		},
		{
			name:     "Parsing coverage file",
			file:     "example1.cov",
			expected: 10,
		},
		{
			name:     "Parsing coverage file, skipping erroneous lines",
			file:     "example2.cov",
			expected: 9,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			file, err := os.Open(path.Join("./testdata", "coverageParser", testcase.file))
			if err != nil {
				t.Fatalf("Unexpected error happened while opening coverage file: %v", err)
			}
			defer file.Close()

			got := len(domain.ParseFile(file))
			if got != testcase.expected {
				t.Fatalf("Expected %d parsed lines, but got %d", testcase.expected, got)
			}
		})
	}
}
