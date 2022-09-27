package statistic

import (
	"io"

	"github.com/engelmi/limgo/pkg"
)

type CoverageStatistic struct {
	GlobalStatistic GlobalStatistic
	FilesStatistics FileStatistics
}

func (cs CoverageStatistic) Print(output io.Writer, verbosity uint) {
	PrintPretty(output, &cs, verbosity)
}

func (cs CoverageStatistic) GetCoveragePercent(directory string) float64 {
	coverage := 0.0
	count := 0
	for _, fs := range cs.FilesStatistics {
		if fs.Dir == directory {
			coverage = coverage + fs.GetCoveragePercent()
			count += 1
		}
	}

	if count == 0 {
		return 0.0
	}
	return coverage / float64(count)
}

type GlobalStatistic struct {
	ModuleName string
	Statistic
}

func (global GlobalStatistic) GetCoveragePercent() float64 {
	if global.Total <= 0 {
		return 100.0
	}
	return float64(global.CountCovered*100) / float64(global.Total)
}

type FileStatistics []FileStatistic

func (fs FileStatistics) GetCoveragePercent() float64 {
	l := float64(len(fs))
	if l <= 0 {
		return 100.0
	}

	covSum := 0.0
	for _, fileStat := range fs {
		covSum = covSum + fileStat.GetCoveragePercent()
	}
	return float64(covSum) / l
}

type FileStatistic struct {
	Name      string
	Dir       string
	FuncStats FuncStatistics
}

func (fs FileStatistic) GetCoveragePercent() float64 {
	countCovered := 0
	total := 0
	for _, funcStats := range fs.FuncStats {
		countCovered = countCovered + funcStats.CountCovered
		total = total + funcStats.Total
	}
	if total <= 0 {
		return 100.0
	}
	return float64(countCovered*100) / float64(total)
}

func (fs FileStatistic) GetUncoveredLines() (uncoveredLines []int) {
	for _, funcStats := range fs.FuncStats {
		uncoveredLines = append(uncoveredLines, funcStats.UncoveredLines...)
	}
	return uncoveredLines
}

type FuncStatistics []FuncStatistic

func (fs FuncStatistics) GetCoveragePercent() float64 {
	covSum := 0.0
	for _, funcStats := range fs {
		covSum = covSum + funcStats.GetCoveragePercent()
	}
	l := float64(len(fs))
	if l <= 0 {
		return 100.0
	}
	return float64(covSum) / l
}

type FuncStatistic struct {
	Name string
	Statistic
}

func (fs FuncStatistic) GetCoveragePercent() float64 {
	if fs.Total <= 0 {
		return 100.0
	}
	return float64(fs.CountCovered*100) / float64(fs.Total)
}

func (fs FuncStatistic) GetUncoveredLines() (uncoveredLines []int) {
	return fs.UncoveredLines
}

type Statistic struct {
	Total          int
	CountCovered   int
	UncoveredLines []int
}

func BuildCoverageStatistic(module *pkg.Module) (*CoverageStatistic, error) {
	fileGroups := module.Coverage.GroupByFile(module.GoMod.ModuleName)
	includedFiles, err := module.IncludedFiles()
	if err != nil {
		return nil, err
	}

	fileStats := []FileStatistic{}
	for _, srcInfo := range includedFiles {
		funcStats := []FuncStatistic{}

		covInfo := fileGroups[srcInfo.Path]

		for _, funcInfo := range srcInfo.FuncsInfo {
			total := 0
			countCovered := 0
			uncoveredLines := []int{}
			for _, stmt := range funcInfo.Stmts {
				total = total + 1
				found := false
				isCovered := false
				for _, line := range covInfo {
					found = stmt.Pos.LineFrom >= line.MustLineFrom() && stmt.Pos.LineFrom <= line.MustLineTo()
					if found {
						isCovered = line.IsCovered()
						break
					}
				}

				if found && isCovered {
					countCovered = countCovered + 1
					continue
				}
				uncoveredLines = append(uncoveredLines, stmt.Pos.LineFrom)
			}
			funcStats = append(funcStats,
				FuncStatistic{
					Name: funcInfo.Name,
					Statistic: Statistic{
						Total:          total,
						CountCovered:   countCovered,
						UncoveredLines: uncoveredLines,
					},
				},
			)
		}
		fileStats = append(fileStats,
			FileStatistic{
				Name:      srcInfo.Name,
				Dir:       srcInfo.Directory(),
				FuncStats: funcStats,
			},
		)
	}

	return &CoverageStatistic{
		GlobalStatistic: GlobalStatistic{
			ModuleName: module.GoMod.ModuleName,
			Statistic: Statistic{
				Total:          includedFiles.NumberOfStatements(),
				CountCovered:   module.Coverage.CountCoveredStatements(),
				UncoveredLines: nil,
			},
		},
		FilesStatistics: fileStats,
	}, nil
}
