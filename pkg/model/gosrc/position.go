package gosrc

import "go/token"

type Position struct {
	LineFrom   int
	LineTo     int
	ColumnFrom int
	ColumnTo   int
}

func NewPosition(from token.Position, to token.Position) Position {
	return Position{
		LineFrom:   from.Line,
		LineTo:     to.Line,
		ColumnFrom: from.Column,
		ColumnTo:   to.Column,
	}
}

func (pos Position) Clone() Position {
	return Position{
		LineFrom:   pos.LineFrom,
		LineTo:     pos.LineTo,
		ColumnFrom: pos.ColumnFrom,
		ColumnTo:   pos.ColumnTo,
	}
}
