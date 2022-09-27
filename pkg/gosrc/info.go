package gosrc

import "strings"

type SourceFilesInfo []SourceFileInfo

func (srcFilesInfo SourceFilesInfo) NumberOfStatements() int {
	sum := 0
	for _, srcFileInfo := range srcFilesInfo {
		for _, funcInfo := range srcFileInfo.FuncsInfo {
			sum = sum + funcInfo.NumberOfStatements()
		}
	}
	return sum
}

type SourceFileInfo struct {
	Name      string
	Path      string
	FuncsInfo FuncsInfo
}

func (srcInfo SourceFileInfo) Directory() string {
	return strings.Replace(srcInfo.Path, "/"+srcInfo.Name, "", 1)
}

type FuncsInfo []FuncInfo

func (funcsInfo FuncsInfo) NumberOfStatements() int {
	sum := 0
	for _, funcInfo := range funcsInfo {
		sum = sum + funcInfo.NumberOfStatements()
	}
	return sum
}

type FuncInfo struct {
	Name  string
	Pos   PosInfo
	Stmts StmtsInfo
}

func (funcInfo FuncInfo) NumberOfStatements() int {
	return len(funcInfo.Stmts)
}

type StmtsInfo []StmtInfo

type StmtInfo struct {
	Pos PosInfo
}

type PosInfo struct {
	PosFrom  int
	PosTo    int
	LineFrom int
	LineTo   int
}
