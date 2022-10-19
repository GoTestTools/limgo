package gosrc

import (
	"go/ast"
	"go/token"
)

type Statement struct {
	Type             ExpressionType
	Position         Position
	ParentStatement  *Statement
	NestedStatements Statements
}

func NewStatement(stmt ast.Stmt, parent *Statement, fs *token.FileSet) Statement {
	from := fs.File(stmt.Pos()).Position(stmt.Pos())
	to := fs.File(stmt.End()).Position(stmt.End())

	newStmt := Statement{
		Type:     ExprTypeFromStmt(stmt),
		Position: NewPosition(from, to),
	}
	if parent != nil {
		newStmt.ParentStatement = parent
	}

	return newStmt
}

func (stmt Statement) IsBranchStmt() bool {
	return stmt.Type == ExprTypeIf ||
		stmt.Type == ExprTypeSwitch ||
		stmt.Type == ExprTypeTypeSwitch ||
		stmt.Type == ExprTypeCaseClause ||
		stmt.Type == ExprTypeFor ||
		stmt.Type == ExprTypeRange
}
