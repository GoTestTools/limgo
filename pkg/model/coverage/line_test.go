package coverage_test

import (
	"testing"

	"github.com/GoTestTools/limgo/pkg/model/coverage"
)

var validLine = coverage.ParsedLine{"github.com/GoTestTools/limgo", "line.go", "11", "5", "14", "58", "3", "1"}
var invalidLine = coverage.ParsedLine{"github.com/GoTestTools/limgo", "line.go", "na", "na", "na", "na", "na", "1"}

func TestParsedLine_Key(t *testing.T) {
	expected := "github.com/GoTestTools/limgo/line.go"
	if validLine.Key() != expected {
		t.Fatalf("Expected line key to be '%s', but got '%s'", expected, validLine.Key())
	}
}

func TestParsedLine_Package(t *testing.T) {
	expected := "github.com/GoTestTools/limgo"
	if validLine.PackageName() != expected {
		t.Fatalf("Expected package to be '%s', but got '%s'", expected, validLine.PackageName())
	}
}

func TestParsedLine_File(t *testing.T) {
	expected := "line.go"
	if validLine.FileName() != expected {
		t.Fatalf("Expected file to be '%s', but got '%s'", expected, validLine.FileName())
	}
}

func TestParsedLine_LineFrom(t *testing.T) {
	expected := 11
	lineFrom, err := validLine.LineFrom()
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v", err)
	}
	if lineFrom != expected {
		t.Fatalf("Expected line from to be '%d', but got '%d'", expected, lineFrom)
	}
}

func TestParsedLine_ColumnFrom(t *testing.T) {
	expected := 5
	columnFrom, err := validLine.ColumnFrom()
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v", err)
	}
	if columnFrom != expected {
		t.Fatalf("Expected column from to be '%d', but got '%d'", expected, columnFrom)
	}
}

func TestParsedLine_LineTo(t *testing.T) {
	expected := 14
	lineTo, err := validLine.LineTo()
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v", err)
	}
	if lineTo != expected {
		t.Fatalf("Expected line to to be '%d', but got '%d'", expected, lineTo)
	}
}

func TestParsedLine_ColumnTo(t *testing.T) {
	expected := 58
	columnTo, err := validLine.ColumnTo()
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v", err)
	}
	if columnTo != expected {
		t.Fatalf("Expected column to to be '%d', but got '%d'", expected, columnTo)
	}
}

func TestParsedLine_NumberOfStatements(t *testing.T) {
	expected := 3
	numberOfStatements, err := validLine.NumberOfStatements()
	if err != nil {
		t.Fatalf("Unexpected error occurred: %v", err)
	}
	if numberOfStatements != expected {
		t.Fatalf("Expected number of statements to be '%d', but got '%d'", expected, numberOfStatements)
	}
}

func TestParsedLine_IsCovered(t *testing.T) {

	testcases := []struct {
		name     string
		line     coverage.ParsedLine
		expected bool
	}{
		{
			name:     "line is not covered",
			line:     coverage.ParsedLine{"github.com/GoTestTools/limgo", "line.go", "11", "5", "14", "58", "3", "0"},
			expected: false,
		},
		{
			name:     "line is covered once",
			line:     coverage.ParsedLine{"github.com/GoTestTools/limgo", "line.go", "11", "5", "14", "58", "3", "1"},
			expected: true,
		},
		{
			name:     "line is covered more than once",
			line:     coverage.ParsedLine{"github.com/GoTestTools/limgo", "line.go", "11", "5", "14", "58", "3", "3"},
			expected: true,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			if testcase.line.IsCovered() != testcase.expected {
				t.Fatalf("Expected line to be covered(%t), but got %t", testcase.expected, validLine.IsCovered())
			}
		})
	}
}

func TestParsedLine_MustLineFromPanics(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Fatalf("Expected must line from to panic, but it did not")
		}
	}()

	invalidLine.MustLineFrom()
}

func TestParsedLine_MustColumnFromPanics(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Fatalf("Expected must column from to panic, but it did not")
		}
	}()

	invalidLine.MustColumnFrom()
}

func TestParsedLine_MustLineToPanics(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Fatalf("Expected must line to to panic, but it did not")
		}
	}()

	invalidLine.MustLineTo()
}

func TestParsedLine_MustColumnToPanics(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Fatalf("Expected must column to to panic, but it did not")
		}
	}()

	invalidLine.MustColumnTo()
}

func TestParsedLine_MustNumberOfStatementsPanics(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Fatalf("Expected must number of statements to panic, but it did not")
		}
	}()

	invalidLine.MustNumberOfStatements()
}
