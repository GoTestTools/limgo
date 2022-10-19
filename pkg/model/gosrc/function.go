package gosrc

import (
	"go/ast"
	"go/token"
)

type Function struct {
	Name       string
	Position   Position
	Statements []Statement
}

func NewFunction(funcDecl ast.FuncDecl, fs *token.FileSet) Function {
	from := fs.File(funcDecl.Body.Lbrace).Position(funcDecl.Body.Lbrace)
	to := fs.File(funcDecl.Body.Rbrace).Position(funcDecl.Body.Rbrace)

	return Function{
		Name: funcDecl.Name.Name,
		Position: Position{
			LineFrom:   from.Line,
			LineTo:     to.Line,
			ColumnFrom: from.Column,
			ColumnTo:   to.Column,
		},
	}
}

func (f Function) GetName() string {
	return f.Name
}

func (f Function) GetStatements() []Statement {
	return f.Statements
}
