package gosrc

import (
	"strings"
)

type AnalyzedFile struct {
	FileName  string
	FilePath  string
	Functions []Function
}

func (file AnalyzedFile) Directory() string {
	return strings.Replace(file.FilePath, "/"+file.FileName, "", 1)
}
