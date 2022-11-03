package domain

import (
	"sort"
	"strings"

	"github.com/GoTestTools/limgo/pkg/model/coverage"
	"github.com/GoTestTools/limgo/pkg/model/gosrc"
	"github.com/GoTestTools/limgo/pkg/model/module"
	"github.com/GoTestTools/limgo/pkg/model/statistic"
)

func BuildStatementStatistic(module module.Module, parsedCoverage coverage.CoverageLines) *statistic.ModuleStatistic {
	coverage := GroupCoverageByDirectory(module.GoMod.ModuleName, parsedCoverage)
	moduleStatistics := statistic.ModuleStatistic{
		Name: module.GoMod.ModuleName,
	}

	fileStatisticsPerDir := map[string]statistic.FilesStatistic{}
	for _, file := range module.Files {

		directory := file.Directory()
		fileStatistics := statistic.FileStatistic{
			Name:      file.FileName,
			Directory: directory,
		}
		for _, function := range file.Functions {

			functionStatistics := statistic.FunctionStatistic{
				Name:     function.GetName(),
				Stmts:    statistic.StatementSummary{},
				Lines:    statistic.LineSummary{},
				Branches: statistic.BranchSummary{},
			}
			for _, stmt := range function.GetStatements() {
				handleCoverage(&functionStatistics, stmt, coverage.IsCovered(file.FilePath, stmt))

				toExploreStmts := stmt.NestedStatements
				for len(toExploreStmts) > 0 {
					elem, _ := toExploreStmts.Pop()
					toExploreStmts.Push(elem.NestedStatements)

					handleCoverage(&functionStatistics, elem, coverage.IsCovered(file.FilePath, elem))
				}
			}
			fileStatistics.FunctionStatistics = append(fileStatistics.FunctionStatistics, functionStatistics)
		}
		fileStatisticsPerDir[directory] = append(fileStatisticsPerDir[directory], fileStatistics)
	}
	moduleStatistics.DirectoryStatistics = buildDirectoryStatistic(fileStatisticsPerDir)

	return &moduleStatistics
}

func handleCoverage(functionStats *statistic.FunctionStatistic, stmt gosrc.Statement, isCovered bool) {
	handleStatement(functionStats, isCovered)
	handleLines(functionStats, stmt, isCovered)
	handleBranch(functionStats, stmt, isCovered)
}

func handleStatement(functionStats *statistic.FunctionStatistic, isCovered bool) {
	functionStats.Stmts.Total += 1
	if isCovered {
		functionStats.Stmts.CountCovered += 1
	}
}

func handleLines(functionStats *statistic.FunctionStatistic, stmt gosrc.Statement, isCovered bool) {
	for i := stmt.Position.LineFrom; i <= stmt.Position.LineTo; i++ {
		functionStats.Lines[i] = isCovered
	}
}

func handleBranch(functionStats *statistic.FunctionStatistic, stmt gosrc.Statement, isCovered bool) {
	if stmt.ParentStatement != nil && stmt.ParentStatement.IsBranchStmt() {
		lineFrom := stmt.ParentStatement.Position.LineFrom
		lineTo := stmt.ParentStatement.Position.LineTo
		if _, exists := functionStats.Branches[lineFrom]; !exists {
			functionStats.Branches[lineFrom] = map[int]bool{}
		}
		functionStats.Branches[lineFrom][lineTo] = isCovered
	}
}

func buildDirectoryStatistic(fileStatsPerDir map[string]statistic.FilesStatistic) (directoryStatistic []statistic.DirectoryStatistic) {
	for directory, fileStatistics := range fileStatsPerDir {
		directoryStatistic = append(directoryStatistic, statistic.DirectoryStatistic{
			Name:           directory,
			FileStatistics: fileStatistics,
		})
	}
	sort.Slice(directoryStatistic, func(i, j int) bool {
		return directoryStatistic[i].Name < directoryStatistic[j].Name
	})
	return directoryStatistic
}

func GroupCoverageByDirectory(moduleName string, coverageLines coverage.CoverageLines) (group coverage.GroupedCoverage) {
	group = coverage.GroupedCoverage{}
	for i := range coverageLines {
		key := strings.Replace(coverageLines[i].Key(), moduleName+"/", "", 1)
		group[key] = append(group[key], coverageLines[i])
	}
	return group
}
