package pkg

import (
	"regexp"

	"github.com/engelmi/limgo/pkg/config"
	"github.com/engelmi/limgo/pkg/coverage"
	"github.com/engelmi/limgo/pkg/gosrc"
)

func GetCurrentModule() (*Module, error) {
	goMod, err := gosrc.GetGoMod()
	if err != nil {
		return nil, err
	}

	files, err := gosrc.GetSourceFiles()
	if err != nil {
		return nil, err
	}

	return &Module{
		GoMod: goMod,
		Files: files,
	}, nil
}

type Module struct {
	GoMod    gosrc.GoMod
	Files    gosrc.SourceFilesInfo
	Coverage *coverage.ParsedLines
	Config   *config.CoverageConfig
}

func (module *Module) WithCoverage(cov coverage.ParsedLines) *Module {
	module.Coverage = &cov
	return module
}

func (module *Module) WithConfig(cfg config.CoverageConfig) *Module {
	module.Config = &cfg
	return module
}

func (module *Module) IncludedFiles() (files gosrc.SourceFilesInfo, err error) {
	shouldInclude, err := createShouldInclude(module)
	if err != nil {
		return nil, err
	}

	for i := range module.Files {
		file := module.Files[i]
		if shouldInclude(file.Path) {
			files = append(files, file)
		}
	}

	return files, nil
}

func createShouldInclude(module *Module) (func(string) bool, error) {
	regs := make([]*regexp.Regexp, 0, len(module.Config.Excludes))
	for _, exclude := range module.Config.Excludes {
		r, err := regexp.Compile(exclude)
		if err != nil {
			return nil, err
		}
		regs = append(regs, r)
	}

	return func(filePath string) bool {
		for _, reg := range regs {
			if reg.MatchString(filePath) {
				return false
			}
		}
		return true
	}, nil
}
