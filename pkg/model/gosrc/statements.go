package gosrc

import "github.com/go-errors/errors"

type Statements []Statement

func (stmts *Statements) Push(newStmts Statements) {
	*stmts = append(*stmts, newStmts...)
}

func (stmts *Statements) Pop() (Statement, error) {
	topIndex := len(*stmts) - 1
	if topIndex < 0 {
		return Statement{}, errors.Errorf("no more statements")
	}
	elem := (*stmts)[topIndex]
	*stmts = (*stmts)[:topIndex]
	return elem, nil
}
