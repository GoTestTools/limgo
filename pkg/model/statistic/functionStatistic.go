package statistic

type FunctionStatistic struct {
	Name     string
	Stmts    StatementSummary
	Lines    LineSummary
	Branches BranchSummary
}

func (stat FunctionStatistic) GetStmtCoverage() float64 {
	return CalculatePercentage(stat.Stmts.Total, stat.Stmts.CountCovered)
}
func (stat FunctionStatistic) GetLinesCoverage() float64 {
	return CalculatePercentage(stat.Lines.GetLinesTotal(), stat.Lines.GetLinesCovered())
}
func (stat FunctionStatistic) GetBranchesCoverage() float64 {
	return CalculatePercentage(stat.Branches.GetBranchesTotal(), stat.Branches.GetBranchesCovered())
}
