package dto

import (
	"fmt"

	"github.com/GoTestTools/limgo/pkg/model/evaluation"
	"github.com/GoTestTools/limgo/pkg/model/statistic"
)

type OutputDTO struct {
	Statistic      *StatisticOutputDTO `json:"statistic,omitempty"`
	CoverageErrors []ErrorOutputDTO    `json:"coverageErrors,omitempty"`
}

func NewOutputDTO(stats *statistic.ModuleStatistic, covErrs []evaluation.CoverageError) OutputDTO {
	statOut := OutputDTOFromStatistic(stats)
	return OutputDTO{
		Statistic:      &statOut,
		CoverageErrors: OutputDTOFromCovErrors(covErrs),
	}
}

type ErrorOutputDTO struct {
	AffectedFile      string `json:"affectedFile"`
	Type              string `json:"type"`
	ExpectedThreshold string `json:"expectedThreshold"`
	ActualCovered     string `json:"actualCovered"`
}

func OutputDTOFromCovErrors(covErrs []evaluation.CoverageError) []ErrorOutputDTO {
	dtos := []ErrorOutputDTO{}

	for _, covErr := range covErrs {
		dtos = append(dtos, ErrorOutputDTO{
			AffectedFile:      covErr.AffectedFile,
			Type:              string(covErr.Type),
			ExpectedThreshold: fmtPercentage(covErr.ExpectedThreshold),
			ActualCovered:     fmtPercentage(covErr.ActualCovered),
		})
	}

	return dtos
}

type StatisticOutputDTO struct {
	Module      string               `json:"module"`
	Coverage    CoverageOutputDTO    `json:"coverage"`
	Directories DirectoriesOutputDTO `json:"directories,omitempty"`
}

type CoverageOutputDTO struct {
	Stmts     string `json:"stmts"`
	Lines     string `json:"lines"`
	Branches  string `json:"branches"`
	Uncovered []int  `json:"uncovered,omitempty"`
}

type DirectoriesOutputDTO map[string]DirectoryOutputDTO

type DirectoryOutputDTO struct {
	Coverage CoverageOutputDTO `json:"coverage"`
	Files    FilesOutputDTO    `json:"files,omitempty"`
}

type FilesOutputDTO map[string]FileOutputDTO

type FileOutputDTO struct {
	Coverage CoverageOutputDTO `json:"coverage"`
	Funcs    FuncsOutputDTO    `json:"funcs,omitempty"`
}

type FuncsOutputDTO map[string]CoverageOutputDTO

func OutputDTOFromStatistic(stats *statistic.ModuleStatistic) StatisticOutputDTO {
	if stats == nil {
		return StatisticOutputDTO{}
	}

	directories := DirectoriesOutputDTO{}
	for _, dirStats := range stats.DirectoryStatistics {

		files := FilesOutputDTO{}
		for _, fileStats := range dirStats.FileStatistics {

			funcs := FuncsOutputDTO{}
			for _, funcStats := range fileStats.FunctionStatistics {
				// exclude GenDecl from the output
				if funcStats.Name == "GenDecl" {
					continue
				}

				funcs[funcStats.Name] = CoverageOutputDTO{
					Stmts:     fmtPercentage(funcStats.GetStmtCoverage()),
					Lines:     fmtPercentage(funcStats.GetLinesCoverage()),
					Branches:  fmtPercentage(funcStats.GetBranchesCoverage()),
					Uncovered: funcStats.Lines.GetUncoveredLines(),
				}
			}

			files[fileStats.Name] = FileOutputDTO{
				Coverage: CoverageOutputDTO{
					Stmts:     fmtPercentage(fileStats.GetStmtCoverage()),
					Lines:     fmtPercentage(fileStats.GetLinesCoverage()),
					Branches:  fmtPercentage(fileStats.GetBranchesCoverage()),
					Uncovered: fileStats.GetUncoveredLines(),
				},
				Funcs: funcs,
			}
		}
		directories[dirStats.Name] = DirectoryOutputDTO{
			Coverage: CoverageOutputDTO{
				Stmts:     fmtPercentage(dirStats.GetStmtCoverage()),
				Lines:     fmtPercentage(dirStats.GetLinesCoverage()),
				Branches:  fmtPercentage(dirStats.GetBranchesCoverage()),
				Uncovered: nil,
			},
			Files: files,
		}
	}

	return StatisticOutputDTO{
		Module: stats.Name,
		Coverage: CoverageOutputDTO{
			Stmts:     fmtPercentage(stats.GetStmtCoverage()),
			Lines:     fmtPercentage(stats.GetLinesCoverage()),
			Branches:  fmtPercentage(stats.GetBranchesCoverage()),
			Uncovered: nil,
		},
		Directories: directories,
	}
}

func fmtPercentage(p float64) string {
	return fmt.Sprintf("%.2f", p)
}
