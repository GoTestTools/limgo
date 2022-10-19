package statistic

import (
	"path"
	"regexp"
)

type ModuleStatistic struct {
	Name                string
	DirectoryStatistics []DirectoryStatistic
}

func (stat ModuleStatistic) GetStmtTotal() (total int) {
	for _, fstat := range stat.DirectoryStatistics {
		total += fstat.GetStmtTotal()
	}
	return total
}

func (stat ModuleStatistic) GetStmtCovered() (covered int) {
	for _, fstat := range stat.DirectoryStatistics {
		covered += fstat.GetStmtCovered()
	}
	return covered
}

func (stat ModuleStatistic) GetStmtCoverage() float64 {
	return CalculatePercentage(stat.GetStmtTotal(), stat.GetStmtCovered())
}

func (stat ModuleStatistic) GetLinesTotal() (total int) {
	for _, fstat := range stat.DirectoryStatistics {
		total += fstat.GetLinesTotal()
	}
	return total
}

func (stat ModuleStatistic) GetLinesCovered() (covered int) {
	for _, fstat := range stat.DirectoryStatistics {
		covered += fstat.GetLinesCovered()
	}
	return covered
}

func (stat ModuleStatistic) GetLinesCoverage() float64 {
	return CalculatePercentage(stat.GetLinesTotal(), stat.GetLinesCovered())
}

func (stat ModuleStatistic) GetBranchesTotal() (total int) {
	for _, fstat := range stat.DirectoryStatistics {
		total += fstat.GetBranchesTotal()
	}
	return total
}

func (stat ModuleStatistic) GetBranchesCovered() (covered int) {
	for _, fstat := range stat.DirectoryStatistics {
		covered += fstat.GetBranchesCovered()
	}
	return covered
}

func (stat ModuleStatistic) GetBranchesCoverage() float64 {
	return CalculatePercentage(stat.GetBranchesTotal(), stat.GetBranchesCovered())
}

func (stat ModuleStatistic) GetDirectoryStatistics() []DirectoryStatistic {
	return stat.DirectoryStatistics
}

func (stat ModuleStatistic) FilterFileStatistics(fileNameFilter *regexp.Regexp) (fileStatistics FilesStatistic) {
	for _, directoryStatistic := range stat.DirectoryStatistics {
		for _, fileStatistic := range directoryStatistic.FileStatistics {
			path := path.Join(fileStatistic.Directory, fileStatistic.Name)
			if fileNameFilter.MatchString(path) {
				fileStatistics = append(fileStatistics, fileStatistic)
			}
		}
	}
	return fileStatistics
}
