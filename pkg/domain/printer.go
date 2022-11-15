package domain

import (
	"fmt"
	"io"

	"github.com/GoTestTools/limgo/pkg/dto"
	"github.com/GoTestTools/limgo/pkg/model/evaluation"
	"github.com/GoTestTools/limgo/pkg/model/statistic"
)

type printer struct {
	w   io.Writer
	fmt dto.Formatter
}

func NewPrinter(w io.Writer, fmt dto.Formatter) printer {
	return printer{
		w:   w,
		fmt: fmt,
	}
}

func (p printer) Print(stats *statistic.ModuleStatistic, errs []evaluation.CoverageError, verbosity uint) {
	output := dto.NewOutputDTO(stats, errs)

	switch verbosity {
	case 0:
		output.Statistic = nil
	case 1:
		output.Statistic.Directories = nil
	case 2:
		for i := range output.Statistic.Directories {
			output.Statistic.Directories[i] = dto.DirectoryOutputDTO{
				Coverage: output.Statistic.Directories[i].Coverage,
			}
		}
	case 3:
		for i := range output.Statistic.Directories {
			for j := range output.Statistic.Directories[i].Files {
				output.Statistic.Directories[i].Files[j] = dto.FileOutputDTO{
					Coverage: output.Statistic.Directories[i].Files[j].Coverage,
				}
			}
		}
	default:
		break
	}

	out, _ := p.fmt(output)
	fmt.Fprint(p.w, out)
}
