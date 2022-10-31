package statistic

import "sort"

type StatementSummary struct {
	Total        int
	CountCovered int
}

// Line key: line number, value: is covered.
type LineSummary map[int]bool

func (ls LineSummary) GetLinesTotal() int {
	return len(ls)
}

func (ls LineSummary) GetLinesCovered() (covered int) {
	for _, isCovered := range ls {
		if isCovered {
			covered += 1
		}
	}
	return covered
}

func (ls LineSummary) GetUncoveredLines() (uncovered []int) {
	for lineNo, isCovered := range ls {
		if !isCovered {
			uncovered = append(uncovered, lineNo)
		}
	}
	sort.Ints(uncovered)
	return uncovered
}

// 1. Key: line number from, 2. Key: line number to, value: is covered.
type BranchSummary map[int]map[int]bool

func (bs BranchSummary) GetBranchesTotal() (total int) {
	for _, to := range bs {
		total += len(to)
	}
	return len(bs)
}

func (bs BranchSummary) GetBranchesCovered() (covered int) {
	for _, to := range bs {
		for _, isCovered := range to {
			if isCovered {
				covered += 1
			}
		}
	}
	return covered
}
