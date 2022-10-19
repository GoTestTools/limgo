package statistic

import "sort"

type FilesStatistic []FileStatistic

func (stat FilesStatistic) GetStmtTotal() (total int) {
	for _, fstat := range stat {
		total += fstat.GetStmtTotal()
	}
	return total
}

func (stat FilesStatistic) GetStmtCovered() (covered int) {
	for _, fstat := range stat {
		covered += fstat.GetStmtCovered()
	}
	return covered
}

func (stat FilesStatistic) GetStmtCoverage() float64 {
	return CalculatePercentage(stat.GetStmtTotal(), stat.GetStmtCovered())
}

func (stat FilesStatistic) GetLinesTotal() (total int) {
	for _, fstat := range stat {
		total += fstat.GetLinesTotal()
	}
	return total
}

func (stat FilesStatistic) GetLinesCovered() (covered int) {
	for _, fstat := range stat {
		covered += fstat.GetLinesCovered()
	}
	return covered
}

func (stat FilesStatistic) GetLinesCoverage() float64 {
	return CalculatePercentage(stat.GetLinesTotal(), stat.GetLinesCovered())
}

func (stat FilesStatistic) GetUncoveredLines() (uncovered []int) {
	for _, fstat := range stat {
		uncovered = append(uncovered, fstat.GetUncoveredLines()...)
	}
	sort.Ints(uncovered)
	return uncovered
}

func (stat FilesStatistic) GetBranchesTotal() (total int) {
	for _, fstat := range stat {
		total += fstat.GetBranchesTotal()
	}
	return total
}

func (stat FilesStatistic) GetBranchesCovered() (covered int) {
	for _, fstat := range stat {
		covered += fstat.GetBranchesCovered()
	}
	return covered
}

func (stat FilesStatistic) GetBranchesCoverage() float64 {
	return CalculatePercentage(stat.GetBranchesTotal(), stat.GetBranchesCovered())
}
