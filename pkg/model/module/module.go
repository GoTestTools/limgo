package module

import (
	"github.com/GoTestTools/limgo/pkg/model/gosrc"
)

type Module struct {
	GoMod GoMod
	Files []gosrc.AnalyzedFile
}
