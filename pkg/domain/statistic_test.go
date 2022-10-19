package domain_test

import (
	"testing"

	"github.com/GoTestTools/limgo/pkg/domain"
	"github.com/GoTestTools/limgo/pkg/model/coverage"
)

func TestGroupCoverageByDirectory(t *testing.T) {
	parsedLines := []coverage.ParsedLine{
		{"github.com/GoTestTools/limgo/pkg/coverage", "line.go", "0", "0", "0", "0", "0", "0"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "line.go", "0", "0", "0", "0", "0", "0"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "parse.go", "0", "0", "0", "0", "0", "0"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "parse.go", "0", "0", "0", "0", "0", "0"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "parse.go", "0", "0", "0", "0", "0", "0"},
		{"github.com/GoTestTools/limgo/pkg/config", "json.go", "0", "0", "0", "0", "0", "0"},
		{"github.com/GoTestTools/limgo/pkg/coverage", "parse.go", "0", "0", "0", "0", "0", "0"},
	}

	grouped := domain.GroupCoverageByDirectory("github.com/GoTestTools/limgo", parsedLines)

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
