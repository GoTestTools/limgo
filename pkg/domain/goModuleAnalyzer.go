package domain

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/GoTestTools/limgo/pkg/model/gosrc"
	"github.com/go-errors/errors"
)

type ModuleAnalyzer func(string) ([]gosrc.AnalyzedFile, error)

var AnalyzeModule = func(rootDir string) (analyzedFiles []gosrc.AnalyzedFile, err error) {
	srcFiles, err := discoverSrcFiles(rootDir)
	if err != nil {
		return nil, err
	}

	for _, goSrcFile := range srcFiles {
		functions, err := exploreFunctions(goSrcFile.Path)
		if err != nil {
			return nil, errors.New(fmt.Errorf("failed opening '%s': %w", goSrcFile, err))
		}
		analyzedFiles = append(analyzedFiles, gosrc.AnalyzedFile{
			FileName:  goSrcFile.Name,
			FilePath:  goSrcFile.Path,
			Functions: functions,
		})
	}
	return analyzedFiles, nil
}

func exploreFunctions(filePath string) (functions []gosrc.Function, err error) {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, filePath, nil, parser.AllErrors)
	if err != nil {
		return nil, errors.New(fmt.Errorf("failed to parse file '%s': %w", filePath, err))
	}

	if f == nil {
		return nil, nil
	}

	for _, decl := range f.Decls {
		switch declType := decl.(type) {
		case *ast.FuncDecl:
			functions = append(functions, exploreFunction(declType, fs))
		case *ast.GenDecl:
			functions = append(functions, exploreGeneralDeclaration(declType, fs))
		default:
		}
	}
	return functions, nil
}

func exploreGeneralDeclaration(genDecl *ast.GenDecl, fs *token.FileSet) gosrc.Function {
	from := fs.File(genDecl.Pos()).Position(genDecl.Pos())
	to := fs.File(genDecl.End()).Position(genDecl.End())

	function := gosrc.Function{
		Name:     string(gosrc.ExprTypeGenDecl),
		Position: gosrc.NewPosition(from, to),
	}

	for _, spec := range genDecl.Specs {
		switch specType := spec.(type) {
		case *ast.ValueSpec:
			for _, value := range specType.Values {
				function.Statements = append(function.Statements, exploreExpression(value, nil, fs)...)
			}
		default:
		}
	}

	return function
}

func exploreFunction(toExplore *ast.FuncDecl, fs *token.FileSet) gosrc.Function {
	from := fs.File(toExplore.Pos()).Position(toExplore.Pos())
	to := fs.File(toExplore.End()).Position(toExplore.End())
	function := gosrc.Function{
		Name:     toExplore.Name.Name,
		Position: gosrc.NewPosition(from, to),
	}

	if toExplore.Body != nil {
		for _, stmt := range toExplore.Body.List {
			s := exploreStatement(stmt, nil, fs)
			function.Statements = append(function.Statements, s)
		}
	}

	return function
}

func exploreStatement(toExplore ast.Stmt, parent *gosrc.Statement, fs *token.FileSet) gosrc.Statement {
	s := gosrc.NewStatement(toExplore, parent, fs)

	switch stmt := toExplore.(type) {
	case *ast.BlockStmt:
		for _, nestedStmt := range stmt.List {
			s.NestedStatements = append(s.NestedStatements, exploreStatement(nestedStmt, &s, fs))
		}
	case *ast.IfStmt:
		ifBody := stmt.Body
		if ifBody != nil {
			for _, nestedStmt := range ifBody.List {
				s.NestedStatements = append(s.NestedStatements, exploreStatement(nestedStmt, &s, fs))
			}
		}
		if stmt.Else != nil {
			s.NestedStatements = append(s.NestedStatements, exploreStatement(stmt.Else, &s, fs))
		}
	case *ast.RangeStmt:
		rangeBody := stmt.Body
		if rangeBody != nil {
			for _, nestedStmt := range rangeBody.List {
				s.NestedStatements = append(s.NestedStatements, exploreStatement(nestedStmt, &s, fs))
			}
		}
	case *ast.ForStmt:
		forBody := stmt.Body
		if forBody != nil {
			for _, nestedStmt := range forBody.List {
				s.NestedStatements = append(s.NestedStatements, exploreStatement(nestedStmt, &s, fs))
			}
		}
	case *ast.TypeSwitchStmt:
		typeSwitchBody := stmt.Body
		if typeSwitchBody != nil {
			for _, nestedStmt := range typeSwitchBody.List {
				s.NestedStatements = append(s.NestedStatements, exploreStatement(nestedStmt, &s, fs))
			}
		}
	case *ast.SwitchStmt:
		switchBody := stmt.Body
		if switchBody != nil {
			for _, nestedStmt := range switchBody.List {
				s.NestedStatements = append(s.NestedStatements, exploreStatement(nestedStmt, &s, fs))
			}
		}
	case *ast.CaseClause:
		caseBody := stmt.Body
		for _, nestedStmt := range caseBody {
			s.NestedStatements = append(s.NestedStatements, exploreStatement(nestedStmt, &s, fs))
		}
	case *ast.AssignStmt:
		for _, rh := range stmt.Rhs {
			s.NestedStatements = append(s.NestedStatements, exploreExpression(rh, &s, fs)...)
		}
	case *ast.DeferStmt:
		s.NestedStatements = append(s.NestedStatements, exploreExpression(stmt.Call.Fun, &s, fs)...)
	case *ast.ExprStmt:
		s.NestedStatements = append(s.NestedStatements, exploreExpression(stmt.X, &s, fs)...)
	case *ast.ReturnStmt:
		for _, result := range stmt.Results {
			s.NestedStatements = append(s.NestedStatements, exploreExpression(result, &s, fs)...)
		}
	default:
	}

	return s
}

func exploreExpression(toExplore ast.Expr, parent *gosrc.Statement, fs *token.FileSet) []gosrc.Statement {
	switch expr := toExplore.(type) {
	case *ast.FuncLit:
		funcBody := expr.Body
		var nestedStatements []gosrc.Statement
		for _, nestedStmt := range funcBody.List {
			nestedStatements = append(nestedStatements, exploreStatement(nestedStmt, parent, fs))
		}
		return nestedStatements
	case *ast.CallExpr:
		args := expr.Args
		var nestedStatements []gosrc.Statement
		for _, arg := range args {
			nestedStatements = append(nestedStatements, exploreExpression(arg, parent, fs)...)
		}
		return nestedStatements
	default:
	}
	return nil
}

type GoSrcFile struct {
	Name string
	Path string
}

func discoverSrcFiles(moduleRootDir string) (srcFiles []GoSrcFile, err error) {
	err = filepath.WalkDir(moduleRootDir, func(relativePath string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// skip vendored files
		if strings.HasPrefix(relativePath, "vendor/") {
			return nil
		}

		if !info.IsDir() && isGoSourceFile(relativePath) {
			srcFiles = append(srcFiles, GoSrcFile{
				Name: info.Name(),
				Path: relativePath,
			})
		}

		return nil
	})
	if err != nil {
		return nil, errors.New(fmt.Errorf("failed to walk over module '%s': %w", moduleRootDir, err))
	}

	return srcFiles, nil
}

func isGoSourceFile(path string) bool {
	return strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, "_test.go")
}
