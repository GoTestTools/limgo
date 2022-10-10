package coverage_test

import (
	"testing"

	"github.com/GoTestTools/limgo/pkg/coverage"
)

func TestFilter(t *testing.T) {
	parsedLines := coverage.ParsedLines{
		{"github.com/GoTestTools/limgo/pkg/coverage", "line.go", "11", "5", "14", "58", "3", "1"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "line.go", "10", "37", "12", "2", "1", "0"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "parse.go", "14", "45", "16", "2", "2", "1"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "parse.go", "18", "42", "20", "2", "2", "0"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "parse.go", "22", "48", "24", "2", "1", "1"},
	}

	filteredLines := parsedLines.Filter(func(line coverage.ParsedLine) bool {
		return line.MustNumberOfStatements() == 2
	})

	if len(parsedLines) != 5 {
		t.Fatalf("Original ParsedLines changed. Expected it to contain %d lines, but has %d", 5, len(parsedLines))
	}
	if len(filteredLines) != 2 {
		t.Fatalf("Expected %d lines matching filter criteria, but got %d lines", 2, len(filteredLines))
	}
}

func TestGroupByFileName(t *testing.T) {
	parsedLines := coverage.ParsedLines{
		{"github.com/GoTestTools/limgo/pkg/coverage", "line.go", "0", "0", "0", "0", "0", "0"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "line.go", "0", "0", "0", "0", "0", "0"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "parse.go", "0", "0", "0", "0", "0", "0"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "parse.go", "0", "0", "0", "0", "0", "0"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "parse.go", "0", "0", "0", "0", "0", "0"},
		{"github.com/GoTestTools/limgo/pkg/config", "json.go", "0", "0", "0", "0", "0", "0"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "parse.go", "0", "0", "0", "0", "0", "0"},
	}

	grouped := parsedLines.GroupByFile("github.com/GoTestTools/limgo")

	if len(grouped) != 3 {
		t.Fatalf("Expected %d groups, but got %d", 3, len(grouped))
	}

	lineGroup := grouped["pkg/coverage/line.go"]
	if len(lineGroup) != 2 {
		t.Fatalf("Expected %d lines in group 'pkg/coverage/line', but got %d", 2, len(lineGroup))
	}
	parseGroup := grouped["pkg/coverage/parse.go"]
	if len(parseGroup) != 4 {
		t.Fatalf("Expected %d lines in group 'pkg/coverage/parse', but got %d", 2, len(parseGroup))
	}
	jsonGroup := grouped["pkg/config/json.go"]
	if len(jsonGroup) != 1 {
		t.Fatalf("Expected %d lines in group 'pkg/config/json', but got %d", 1, len(jsonGroup))
	}
}

func TestNumberOfStatements(t *testing.T) {
	parsedLines := coverage.ParsedLines{
		{"github.com/GoTestTools/limgo/pkg/coverage", "line.go", "11", "5", "14", "58", "3", "1"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "line.go", "10", "37", "12", "2", "1", "0"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "parse.go", "14", "45", "16", "2", "2", "1"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "parse.go", "18", "42", "20", "2", "2", "0"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "parse.go", "22", "48", "24", "2", "1", "1"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "parse.go", "12", "9", "38", "8", "4", "1"},
	}

	total, covered, uncovered := parsedLines.NumberOfStatements()
	if total != 13 {
		t.Errorf("Expeced total number of statements to be %d, but got %d", 13, total)
	}
	if covered != 10 {
		t.Errorf("Expeced covered number of statements to be %d, but got %d", 10, covered)
	}
	if uncovered != 3 {
		t.Errorf("Expeced uncovered number of statements to be %d, but got %d", 3, uncovered)
	}
}

func TestCountOverallStatements(t *testing.T) {
	parsedLines := coverage.ParsedLines{
		{"github.com/GoTestTools/limgo/pkg/coverage", "line.go", "11", "5", "14", "58", "3", "1"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "line.go", "10", "37", "12", "2", "1", "0"},
	}

	total := parsedLines.CountOverallStatements()
	if total != 4 {
		t.Errorf("Expeced total number of statements to be %d, but got %d", 4, total)
	}
}

func TestCountCoveredStatements(t *testing.T) {
	parsedLines := coverage.ParsedLines{
		{"github.com/GoTestTools/limgo/pkg/coverage", "line.go", "11", "5", "14", "58", "3", "1"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "line.go", "10", "37", "12", "2", "1", "0"},
	}

	covered := parsedLines.CountCoveredStatements()
	if covered != 3 {
		t.Errorf("Expeced total number of statements to be %d, but got %d", 3, covered)
	}
}

func TestCountUncoveredStatements(t *testing.T) {
	parsedLines := coverage.ParsedLines{
		{"github.com/GoTestTools/limgo/pkg/coverage", "line.go", "11", "5", "14", "58", "3", "1"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "line.go", "10", "37", "12", "2", "1", "0"},
	}

	uncovered := parsedLines.CountUncoveredStatements()
	if uncovered != 1 {
		t.Errorf("Expeced total number of statements to be %d, but got %d", 1, uncovered)
	}
}

func TestCoveredStatementsPercent(t *testing.T) {
	parsedLines := coverage.ParsedLines{
		{"github.com/GoTestTools/limgo/pkg/coverage", "line.go", "11", "5", "14", "58", "3", "1"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "line.go", "10", "37", "12", "2", "1", "0"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "parse.go", "14", "45", "16", "2", "2", "1"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "parse.go", "18", "42", "20", "2", "2", "0"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "parse.go", "22", "48", "24", "2", "1", "1"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "parse.go", "12", "9", "38", "8", "4", "1"},
	}

	gotPercentageCovered := parsedLines.CoveredStatementsPercent()
	if gotPercentageCovered != 76.92 {
		t.Errorf("Expected percentage of covered statements to be %.2f%%, but got %.2f%%", 76.92, gotPercentageCovered)
	}

	gotPercentageUncovered := parsedLines.UncoveredStatementsPercent()
	if gotPercentageUncovered != 23.08 {
		t.Errorf("Expected percentage of uncovered statements to be %.2f%%, but got %.2f%%", 23.08, gotPercentageUncovered)
	}
}

func TestNumberOfLines(t *testing.T) {
	parsedLines := coverage.ParsedLines{
		{"github.com/GoTestTools/limgo/pkg/coverage", "line.go", "11", "5", "14", "58", "3", "1"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "line.go", "10", "37", "12", "2", "1", "0"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "parse.go", "14", "45", "16", "2", "2", "1"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "parse.go", "18", "42", "20", "2", "2", "0"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "parse.go", "22", "48", "24", "2", "1", "1"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "parse.go", "12", "9", "38", "8", "4", "1"},
	}

	total, covered, uncovered := parsedLines.NumberOfLines()
	if total != 37 {
		t.Errorf("Expeced total number of lines to be %d, but got %d", 37, total)
	}
	if covered != 33 {
		t.Errorf("Expeced covered number of lines to be %d, but got %d", 33, covered)
	}
	if uncovered != 4 {
		t.Errorf("Expeced uncovered number of lines to be %d, but got %d", 4, uncovered)
	}
}

func TestCountOverallLines(t *testing.T) {
	parsedLines := coverage.ParsedLines{
		{"github.com/GoTestTools/limgo/pkg/coverage", "line.go", "11", "5", "14", "58", "3", "1"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "line.go", "10", "37", "12", "2", "1", "0"},
	}

	total := parsedLines.CountOverallLines()
	if total != 5 {
		t.Errorf("Expeced total number of lines to be %d, but got %d", 5, total)
	}
}

func TestCountCoveredLines(t *testing.T) {
	parsedLines := coverage.ParsedLines{
		{"github.com/GoTestTools/limgo/pkg/coverage", "line.go", "11", "5", "14", "58", "3", "1"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "line.go", "10", "37", "12", "2", "1", "0"},
	}

	covered := parsedLines.CountCoveredLines()
	if covered != 3 {
		t.Errorf("Expeced total number of lines to be %d, but got %d", 3, covered)
	}
}

func TestCountUncoveredLines(t *testing.T) {
	parsedLines := coverage.ParsedLines{
		{"github.com/GoTestTools/limgo/pkg/coverage", "line.go", "11", "5", "14", "58", "3", "1"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "line.go", "10", "37", "12", "2", "1", "0"},
	}

	uncovered := parsedLines.CountUncoveredLines()
	if uncovered != 2 {
		t.Errorf("Expeced total number of statements to be %d, but got %d", 2, uncovered)
	}
}

func TestCoveredLinesPercent(t *testing.T) {
	parsedLines := coverage.ParsedLines{
		{"github.com/GoTestTools/limgo/pkg/coverage", "line.go", "11", "5", "14", "58", "3", "1"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "line.go", "10", "37", "12", "2", "1", "0"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "parse.go", "14", "45", "16", "2", "2", "1"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "parse.go", "18", "42", "20", "2", "2", "0"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "parse.go", "22", "48", "24", "2", "1", "1"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "parse.go", "12", "9", "38", "8", "4", "1"},
	}

	gotPercentageCovered := parsedLines.CoveredLinesPercent()
	if gotPercentageCovered != 89.19 {
		t.Errorf("Expected percentage of covered lines to be %.2f%%, but got %.2f%%", 89.19, gotPercentageCovered)
	}

	gotPercentageUncovered := parsedLines.UncoveredLinesPercent()
	if gotPercentageUncovered != 10.81 {
		t.Errorf("Expected percentage of uncovered lines to be %.2f%%, but got %.2f%%", 10.81, gotPercentageUncovered)
	}
}
