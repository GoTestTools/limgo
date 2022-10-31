package domain

import (
	"fmt"
	"io"
	"regexp"

	"github.com/GoTestTools/limgo/pkg/model/module"
	"github.com/go-errors/errors"
)

type GoModParser func(io.Reader) (module.GoMod, error)

const goModRegex string = `module\s(.+)\n\ngo\s(\d\.\d\d)`

var ParseGoMod = func(f io.Reader) (module.GoMod, error) {
	goModContent, err := io.ReadAll(f)
	if err != nil {
		return module.GoMod{}, errors.New(fmt.Errorf("failed reading go.mod: %w", err))
	}

	regex := regexp.MustCompile(goModRegex)
	matches := regex.FindStringSubmatch(string(goModContent))
	if len(matches) != 3 {
		return module.GoMod{}, errors.New("couldn't detect go.mod components")
	}

	return module.GoMod{
		ModuleName: matches[1],
		GoVersion:  matches[2],
	}, nil
}
