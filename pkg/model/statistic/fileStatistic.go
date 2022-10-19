package statistic

import "sort"

type FileStatistic struct {
	Name               string
	Directory          string
	FunctionStatistics []FunctionStatistic
}

func (stat FileStatistic) GetStmtTotal() (total int) {
	for _, fstat := range stat.FunctionStatistics {
		total += fstat.Stmts.Total
	}
	return total
}

func (stat FileStatistic) GetStmtCovered() (covered int) {
	for _, fstat := range stat.FunctionStatistics {
		covered += fstat.Stmts.CountCovered
	}
	return covered
}

func (stat FileStatistic) GetStmtCoverage() float64 {
	return CalculatePercentage(stat.GetStmtTotal(), stat.GetStmtCovered())
}

func (stat FileStatistic) GetLinesTotal() (total int) {
	for _, fstat := range stat.FunctionStatistics {
		total += fstat.Lines.GetLinesTotal()
	}
	return total
}

func (stat FileStatistic) GetLinesCovered() (covered int) {
	for _, fstat := range stat.FunctionStatistics {
		covered += fstat.Lines.GetLinesCovered()
	}
	return covered
}

func (stat FileStatistic) GetLinesCoverage() float64 {
	return CalculatePercentage(stat.GetLinesTotal(), stat.GetLinesCovered())
}

func (stat FileStatistic) GetUncoveredLines() (uncovered []int) {
	for _, fstat := range stat.FunctionStatistics {
		uncovered = append(uncovered, fstat.Lines.GetUncoveredLines()...)
	}
	sort.Ints(uncovered)
	return uncovered
}

func (stat FileStatistic) GetBranchesTotal() (total int) {
	for _, fstat := range stat.FunctionStatistics {
		total += fstat.Branches.GetBranchesTotal()
	}
	return total
}

func (stat FileStatistic) GetBranchesCovered() (covered int) {
	for _, fstat := range stat.FunctionStatistics {
		covered += fstat.Branches.GetBranchesCovered()
	}
	return covered
}

func (stat FileStatistic) GetBranchesCoverage() float64 {
	return CalculatePercentage(stat.GetBranchesTotal(), stat.GetBranchesCovered())
}
