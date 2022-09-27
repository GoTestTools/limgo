package gosrc

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/go-errors/errors"
)

func AnalyzeSourceFile(filePath string) (FuncsInfo, error) {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, filePath, nil, parser.AllErrors)
	if err != nil {
		return nil, errors.New(fmt.Errorf("failed to parse file '%s': %w", filePath, err))
	}

	if f == nil {
		return nil, nil
	}

	return exploreFunctions(f, fs), nil
}

func exploreFunctions(astFile *ast.File, fs *token.FileSet) (funcs FuncsInfo) {
	for _, decl := range astFile.Decls {
		switch t := decl.(type) {
		case *ast.FuncDecl:
			if t.Body != nil {
				stmts := exploreStatements(t.Body.List, fs)

				funcInfo := FuncInfo{
					Name: t.Name.Name,
					Pos: PosInfo{
						PosFrom:  int(t.Body.Lbrace),
						PosTo:    int(t.Body.Rbrace),
						LineFrom: fs.File(t.Body.Lbrace).Position(t.Body.Lbrace).Line,
						LineTo:   fs.File(t.Body.Rbrace).Position(t.Body.Rbrace).Line,
					},
					Stmts: stmts,
				}
				funcs = append(funcs, funcInfo)
			}
		}
	}
	return funcs
}

func exploreStatements(toExplore []ast.Stmt, fs *token.FileSet) StmtsInfo {
	stmts := StmtsInfo{}
	for _, stmt := range toExplore {
		switch t := stmt.(type) {
		case *ast.BlockStmt:
			stmts = append(stmts, exploreStatements(t.List, fs)...)
		case *ast.IfStmt:
			if t.Body != nil {
				stmts = append(stmts, StmtInfo{
					Pos: PosInfo{
						PosFrom:  int(t.Body.Lbrace),
						PosTo:    int(t.Body.Rbrace),
						LineFrom: fs.File(t.Body.Lbrace).Position(t.Body.Lbrace).Line,
						LineTo:   fs.File(t.Body.Rbrace).Position(t.Body.Rbrace).Line,
					},
				})
				stmts = append(stmts, exploreStatements(t.Body.List, fs)...)
			}

			switch isIf := t.Else.(type) {
			case *ast.IfStmt:
				if isIf.Body != nil {
					stmts = append(stmts, StmtInfo{
						Pos: PosInfo{
							PosFrom:  int(isIf.Body.Lbrace),
							PosTo:    int(isIf.Body.Rbrace),
							LineFrom: fs.File(isIf.Body.Lbrace).Position(isIf.Body.Lbrace).Line,
							LineTo:   fs.File(isIf.Body.Rbrace).Position(isIf.Body.Rbrace).Line,
						},
					})
					stmts = append(stmts, exploreStatements(isIf.Body.List, fs)...)
				}
				switch isElse := isIf.Else.(type) {
				case *ast.BlockStmt:
					stmts = append(stmts, StmtInfo{
						Pos: PosInfo{
							PosFrom:  int(isElse.Lbrace),
							PosTo:    int(isElse.Rbrace),
							LineFrom: fs.File(isElse.Lbrace).Position(isElse.Lbrace).Line,
							LineTo:   fs.File(isElse.Rbrace).Position(isElse.Rbrace).Line,
						},
					})
					stmts = append(stmts, exploreStatements(isElse.List, fs)...)
				}
			case *ast.BlockStmt:
				stmts = append(stmts, exploreStatements(isIf.List, fs)...)
			}
		case *ast.RangeStmt:
			if t.Body != nil {
				stmts = append(stmts, StmtInfo{
					Pos: PosInfo{
						PosFrom:  int(t.Body.Lbrace),
						PosTo:    int(t.Body.Rbrace),
						LineFrom: fs.File(t.Body.Lbrace).Position(t.Body.Lbrace).Line,
						LineTo:   fs.File(t.Body.Rbrace).Position(t.Body.Rbrace).Line,
					},
				})
				stmts = append(stmts, exploreStatements(t.Body.List, fs)...)
			}
		case *ast.ForStmt:
			if t.Body != nil {
				stmts = append(stmts, StmtInfo{
					Pos: PosInfo{
						PosFrom:  int(t.Body.Lbrace),
						PosTo:    int(t.Body.Rbrace),
						LineFrom: fs.File(t.Body.Lbrace).Position(t.Body.Lbrace).Line,
						LineTo:   fs.File(t.Body.Rbrace).Position(t.Body.Rbrace).Line,
					},
				})
				stmts = append(stmts, exploreStatements(t.Body.List, fs)...)
			}
		case *ast.SwitchStmt:
			if t.Body != nil {
				stmts = append(stmts, StmtInfo{
					Pos: PosInfo{
						PosFrom:  int(t.Body.Lbrace),
						PosTo:    int(t.Body.Rbrace),
						LineFrom: fs.File(t.Body.Lbrace).Position(t.Body.Lbrace).Line,
						LineTo:   fs.File(t.Body.Rbrace).Position(t.Body.Rbrace).Line,
					},
				})
				stmts = append(stmts, exploreStatements(t.Body.List, fs)...)
			}
		case *ast.CaseClause:
			stmts = append(stmts, exploreStatements(t.Body, fs)...)
		case *ast.ReturnStmt:
			stmts = append(stmts, StmtInfo{
				Pos: PosInfo{
					PosFrom:  int(t.Pos()),
					PosTo:    int(t.End()),
					LineFrom: fs.File(t.Pos()).Position(t.Pos()).Line,
					LineTo:   fs.File(t.End()).Position(t.End()).Line,
				},
			})
		default:
			stmts = append(stmts, StmtInfo{
				Pos: PosInfo{
					PosFrom:  int(t.Pos()),
					PosTo:    int(t.End()),
					LineFrom: fs.File(t.Pos()).Position(t.Pos()).Line,
					LineTo:   fs.File(t.End()).Position(t.End()).Line,
				},
			})
		}
	}
	return stmts
}
