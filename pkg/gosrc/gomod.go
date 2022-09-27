package gosrc

import (
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/go-errors/errors"
)

const GoModRegex string = `module\s(.+)\n\ngo\s(\d\.\d\d)`

type GoMod struct {
	ModuleName string
	GoVersion  string
}

func GetGoMod() (GoMod, error) {
	f, err := os.Open("./go.mod")
	if err != nil {
		return GoMod{}, errors.New(fmt.Errorf("failed opening go.mod: %w", err))
	}
	defer f.Close()

	goModContent, err := io.ReadAll(f)
	if err != nil {
		return GoMod{}, errors.New(fmt.Errorf("failed reading go.mod: %w", err))
	}

	regex, err := regexp.Compile(GoModRegex)
	if err != nil {
		return GoMod{}, errors.New(fmt.Errorf("invalid go.mod regex: %w", err))
	}

	matches := regex.FindStringSubmatch(string(goModContent))
	if len(matches) != 3 {
		return GoMod{}, errors.New("couldn't detect go.mod components")
	}

	return GoMod{
		ModuleName: matches[1],
		GoVersion:  matches[2],
	}, nil
}
