package coverage

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type ParsedLines []ParsedLine

func (lines ParsedLines) Filter(f func(line ParsedLine) bool) ParsedLines {
	var applicableLines ParsedLines
	for i := range lines {
		if f(lines[i]) {
			applicableLines = append(applicableLines, lines[i])
		}
	}
	return applicableLines
}

func (lines ParsedLines) GroupByFile(prefixToRemove string) (group map[string]ParsedLines) {
	group = map[string]ParsedLines{}
	for i := range lines {
		key := strings.Replace(lines[i].Key(), prefixToRemove+"/", "", 1)
		group[key] = append(group[key], lines[i])
	}
	return group
}

func (lines ParsedLines) NumberOfStatements() (int, int, int) {
	overall := 0
	covered := 0
	uncovered := 0

	for _, line := range lines {
		count, _ := line.NumberOfStatements()
		overall = overall + count
		if line.IsCovered() {
			covered += count
			continue
		}
		uncovered += count
	}

	return overall, covered, uncovered
}

func (lines ParsedLines) CountOverallStatements() int {
	overall, _, _ := lines.NumberOfStatements()
	return overall
}

func (lines ParsedLines) CountCoveredStatements() int {
	_, covered, _ := lines.NumberOfStatements()
	return covered
}

func (lines ParsedLines) CountUncoveredStatements() int {
	_, _, uncovered := lines.NumberOfStatements()
	return uncovered
}

func (lines ParsedLines) CoveredStatementsPercent() float64 {
	overall, covered, _ := lines.NumberOfStatements()

	if overall <= 0 {
		return 0
	}

	percentage := (float64(covered) / float64(overall)) * 100
	formattedPercentage := fmt.Sprintf("%.2f", percentage)

	value, _ := strconv.ParseFloat(formattedPercentage, 64)
	return math.Floor(value*100) / 100
}

func (lines ParsedLines) UncoveredStatementsPercent() float64 {
	return math.Floor(
		(100-lines.CoveredStatementsPercent())*100,
	) / 100
}

func (lines ParsedLines) NumberOfLines() (int, int, int) {
	overall := 0
	covered := 0
	uncovered := 0

	for _, line := range lines {
		to, _ := line.LineTo()
		from, _ := line.LineFrom()
		count := to - from

		overall = overall + count
		if line.IsCovered() {
			covered += count
			continue
		}
		uncovered += count
	}

	return overall, covered, uncovered
}

func (lines ParsedLines) CountOverallLines() int {
	overall, _, _ := lines.NumberOfLines()
	return overall
}

func (lines ParsedLines) CountCoveredLines() int {
	_, covered, _ := lines.NumberOfLines()
	return covered
}

func (lines ParsedLines) CountUncoveredLines() int {
	_, _, uncovered := lines.NumberOfLines()
	return uncovered
}

func (lines ParsedLines) CoveredLinesPercent() float64 {
	overall, covered, _ := lines.NumberOfLines()

	if overall <= 0 {
		return 0
	}

	percentage := (float64(covered) / float64(overall)) * 100
	formattedPercentage := fmt.Sprintf("%.2f", percentage)

	value, _ := strconv.ParseFloat(formattedPercentage, 64)
	return math.Floor(value*100) / 100
}

func (lines ParsedLines) UncoveredLinesPercent() float64 {
	return math.Floor(
		(100-lines.CoveredLinesPercent())*100,
	) / 100
}
