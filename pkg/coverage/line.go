package coverage

import (
	"path"
	"strconv"
)

type ParsedLine []string

func (line ParsedLine) Key() string {
	return path.Join(line.PackageName(), line.FileName())
}

func (line ParsedLine) PackageName() string {
	return line[0]
}

func (line ParsedLine) FileName() string {
	return line[1]
}

func (line ParsedLine) LineFrom() (int, error) {
	return strconv.Atoi(line[2])
}

func (line ParsedLine) ColumnFrom() (int, error) {
	return strconv.Atoi(line[3])
}

func (line ParsedLine) LineTo() (int, error) {
	return strconv.Atoi(line[4])
}

func (line ParsedLine) ColumnTo() (int, error) {
	return strconv.Atoi(line[5])
}

func (line ParsedLine) NumberOfStatements() (int, error) {
	return strconv.Atoi(line[6])
}

func (line ParsedLine) IsCovered() bool {
	// 0 == statements have not been covered
	// n > 0 == statements have been covered n times
	return line[7] != "0"
}

func (line ParsedLine) MustLineFrom() int {
	v, err := line.LineFrom()
	if err != nil {
		panic(err)
	}
	return v
}

func (line ParsedLine) MustColumnFrom() int {
	v, err := line.ColumnFrom()
	if err != nil {
		panic(err)
	}
	return v
}

func (line ParsedLine) MustLineTo() int {
	v, err := line.LineTo()
	if err != nil {
		panic(err)
	}
	return v
}

func (line ParsedLine) MustColumnTo() int {
	v, err := line.ColumnTo()
	if err != nil {
		panic(err)
	}
	return v
}

func (line ParsedLine) MustNumberOfStatements() int {
	v, err := line.NumberOfStatements()
	if err != nil {
		panic(err)
	}
	return v
}
