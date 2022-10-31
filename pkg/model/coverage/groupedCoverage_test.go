package coverage_test

import (
	"testing"

	"github.com/GoTestTools/limgo/pkg/model/coverage"
	"github.com/GoTestTools/limgo/pkg/model/gosrc"
)

func TestIsCovered(t *testing.T) {

	groupedCoverage := coverage.GroupedCoverage{}
	groupedCoverage["pkg/dto/json.go"] = []coverage.ParsedLine{
		{"github.com/GoTestTools/limgo/pkg/dto", "json.go", "8", "56", "12", "16", "4", "1"},
		{"github.com/GoTestTools/limgo/pkg/dto", "json.go", "15", "2", "15", "12", "1", "1"},
		{"github.com/GoTestTools/limgo/pkg/dto", "json.go", "12", "16", "14", "3", "1", "0"},
		{"github.com/GoTestTools/limgo/pkg/dto", "json.go", "18", "58", "22", "16", "4", "1"},
		{"github.com/GoTestTools/limgo/pkg/dto", "json.go", "25", "2", "25", "20", "1", "1"},
		{"github.com/GoTestTools/limgo/pkg/dto", "json.go", "22", "16", "24", "3", "1", "1"},
	}

	testcases := []struct {
		name     string
		fileName string
		stmt     gosrc.Statement
		expected bool
	}{
		{
			name:     "returns false if file is not listed in coverage",
			fileName: "donotexist.go",
			expected: false,
		},
		{
			name:     "returns false if statement is not listed in coverage",
			fileName: "pkg/dto/json.go",
			stmt: gosrc.Statement{Position: gosrc.Position{
				LineFrom: 1, LineTo: 1,
			}},
			expected: false,
		},
		{
			name:     "returns true if statement is covered and analyzed stmt starts exaclty as listed in coverage",
			fileName: "pkg/dto/json.go",
			stmt: gosrc.Statement{Position: gosrc.Position{
				LineFrom: 8, LineTo: 11,
			}},
			expected: true,
		},
		{
			name:     "returns true if statement is covered and analyzed stmt starts within in coverage range",
			fileName: "pkg/dto/json.go",
			stmt: gosrc.Statement{Position: gosrc.Position{
				LineFrom: 9, LineTo: 11,
			}},
			expected: true,
		},
		{
			name:     "returns false if statement is not covered and analyzed stmt starts within in coverage range",
			fileName: "pkg/dto/json.go",
			stmt: gosrc.Statement{Position: gosrc.Position{
				LineFrom: 13, LineTo: 13,
			}},
			expected: false,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			got := groupedCoverage.IsCovered(testcase.fileName, testcase.stmt)
			if testcase.expected != got {
				t.Fatalf("Expected statement at position %v to be covered (%t)", testcase.stmt.Position, testcase.expected)
			}
		})
	}
}
