package domain

import (
	"bufio"
	"io"
	"strings"

	"github.com/GoTestTools/limgo/pkg/model/coverage"
)

type Parser func(io.Reader) coverage.CoverageLines

var ParseFile = func(file io.Reader) coverage.CoverageLines {
	lineMatcher := coverage.NewLineMatcher()
	parsedLines := coverage.CoverageLines{}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "mode:") {
			continue
		}

		parsedLine, err := lineMatcher.Apply(line)
		if err == nil {
			parsedLines = append(parsedLines, parsedLine)
		}
	}

	return parsedLines
}
