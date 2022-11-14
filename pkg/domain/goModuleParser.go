package domain

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/GoTestTools/limgo/pkg/model/module"
	"github.com/go-errors/errors"
)

type ModuleParser func(string, ModuleAnalyzer) (module.Module, error)

var ParseModule = func(rootDir string, srcAnalyzer ModuleAnalyzer) (module.Module, error) {
	gomodFile, err := os.Open(filepath.Clean(path.Join(rootDir, "go.mod")))
	if err != nil {
		return module.Module{}, errors.New(fmt.Errorf("failed opening go.mod: %w", err))
	}
	//nolint:errcheck,gosec
	defer gomodFile.Close()

	gomod, err := ParseGoMod(gomodFile)
	if err != nil {
		return module.Module{}, err
	}

	analyzedFiles, err := srcAnalyzer(rootDir)
	if err != nil {
		return module.Module{}, err
	}

	return module.Module{
		GoMod: gomod,
		Files: analyzedFiles,
	}, nil
}
