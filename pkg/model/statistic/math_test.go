package statistic_test

import (
	"testing"

	"github.com/GoTestTools/limgo/pkg/model/statistic"
)

func TestCalculatePercentage(t *testing.T) {
	testcases := []struct {
		name     string
		total    int
		partial  int
		expected float64
	}{
		{
			name:     "returns 0 if total is 0",
			total:    0,
			partial:  90,
			expected: 0.0,
		},
		{
			name:     "returns float with 2 digit precision",
			total:    11,
			partial:  3,
			expected: 27.27,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			got := statistic.CalculatePercentage(testcase.total, testcase.partial)
			if testcase.expected != got {
				t.Fatalf("Exptected calculated percentage of %v, but got %v", testcase.expected, got)
			}
		})
	}
}
