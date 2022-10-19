package statistic

type DirectoryStatistic struct {
	Name           string
	FileStatistics FilesStatistic
}

func (stat DirectoryStatistic) GetName() string {
	return stat.Name
}

func (stat DirectoryStatistic) GetStmtTotal() int {
	return stat.FileStatistics.GetStmtTotal()
}

func (stat DirectoryStatistic) GetStmtCovered() int {
	return stat.FileStatistics.GetStmtCovered()
}

func (stat DirectoryStatistic) GetStmtCoverage() float64 {
	return CalculatePercentage(stat.GetStmtTotal(), stat.GetStmtCovered())
}

func (stat DirectoryStatistic) GetLinesTotal() int {
	return stat.FileStatistics.GetLinesTotal()
}

func (stat DirectoryStatistic) GetLinesCovered() int {
	return stat.FileStatistics.GetLinesCovered()
}

func (stat DirectoryStatistic) GetLinesCoverage() float64 {
	return CalculatePercentage(stat.GetLinesTotal(), stat.GetLinesCovered())
}

func (stat DirectoryStatistic) GetBranchesTotal() int {
	return stat.FileStatistics.GetBranchesTotal()
}

func (stat DirectoryStatistic) GetBranchesCovered() int {
	return stat.FileStatistics.GetBranchesCovered()
}

func (stat DirectoryStatistic) GetBranchesCoverage() float64 {
	return CalculatePercentage(stat.GetBranchesTotal(), stat.GetBranchesCovered())
}

func (stat DirectoryStatistic) GetFileStatistics() FilesStatistic {
	return stat.FileStatistics
}
