package gosrc

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func GetSourceFiles() (srcs SourceFilesInfo, err error) {
	err = filepath.WalkDir(".", func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && isGoSourceFile(path) {

			funcs, err := AnalyzeSourceFile(path)
			if err != nil {
				return err
			}

			srcs = append(srcs, SourceFileInfo{
				Name:      info.Name(),
				Path:      path,
				FuncsInfo: funcs,
			})
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return srcs, nil
}

func isGoSourceFile(path string) bool {
	return strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, "_test.go")
}
