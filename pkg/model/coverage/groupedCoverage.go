package coverage

import (
	"github.com/GoTestTools/limgo/pkg/model/gosrc"
)

type GroupedCoverage map[string][]ParsedLine

func (pc GroupedCoverage) IsCovered(file string, stmt gosrc.Statement) bool {
	linesOfFile := pc[file]
	if len(linesOfFile) == 0 {
		return false
	}

	for _, line := range linesOfFile {
		if stmt.Position.LineFrom >= line.MustLineFrom() && stmt.Position.LineFrom <= line.MustLineTo() {
			return line.IsCovered()
		}
	}
	return false
}
