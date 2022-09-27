package coverage

import (
	"bufio"
	"io"
	"strings"
)

func Parse(file io.Reader) (ParsedLines, error) {
	lineMatcher := NewLineMatcher()
	parsedLines := ParsedLines{}

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

	return parsedLines, nil
}
