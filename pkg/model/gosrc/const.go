package gosrc

import (
	"fmt"
	"go/ast"
	"strings"
)

type ExpressionType string

func ExprTypeFromStmt(stmt ast.Stmt) ExpressionType {
	return ExpressionType(strings.TrimPrefix(fmt.Sprintf("%T", stmt), "*ast."))
}

const (
	ExprTypeBlock       ExpressionType = "BlockStmt"
	ExprTypeIf          ExpressionType = "IfStmt"
	ExprTypeSwitch      ExpressionType = "SwitchStmt"
	ExprTypeTypeSwitch  ExpressionType = "TypeSwitchStmt"
	ExprTypeCaseClause  ExpressionType = "CaseClause"
	ExprTypeRange       ExpressionType = "RangeStmt"
	ExprTypeFor         ExpressionType = "ForStmt"
	ExprTypeAssign      ExpressionType = "AssignStmt"
	ExprTypeDeclaration ExpressionType = "DeclStmt"
	ExprTypeDefer       ExpressionType = "DeferStmt"
	ExprTypeExpr        ExpressionType = "ExprStmt"

	ExprTypeGenDecl ExpressionType = "GenDecl"
)
