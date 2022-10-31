package statistic_test

import (
	"testing"

	"github.com/GoTestTools/limgo/pkg/model/statistic"
	"github.com/google/go-cmp/cmp"
)

func TestLineSummaryGetLinesTotal(t *testing.T) {
	summary := statistic.LineSummary{}
	summary[1] = true
	summary[2] = false
	summary[3] = true
	summary[5] = true
	if summary.GetLinesTotal() != 4 {
		t.Fatalf("Expected summary lines total to be %d, but got %d", 4, summary.GetLinesTotal())
	}
}

func TestLineSummaryGetLinesCovered(t *testing.T) {
	summary := statistic.LineSummary{}
	summary[1] = true
	summary[2] = false
	summary[3] = true
	summary[5] = true
	if summary.GetLinesCovered() != 3 {
		t.Fatalf("Expected summary lines total to be %d, but got %d", 3, summary.GetLinesCovered())
	}
}

func TestLineSummaryGetUncoveredLines(t *testing.T) {
	summary := statistic.LineSummary{}
	summary[1] = true
	summary[2] = false
	summary[3] = true
	summary[5] = true
	summary[11] = false
	summary[4] = false
	expected := []int{2, 4, 11}
	if diff := cmp.Diff(expected, summary.GetUncoveredLines()); diff != "" {
		t.Fatalf("Detected diff in sorted, uncovered lines: %s", diff)
	}
}

func TestBranchSummaryGetBranchesTotal(t *testing.T) {
	summary := statistic.BranchSummary{}
	summary[9] = map[int]bool{}
	summary[9][11] = true
	summary[14] = map[int]bool{}
	summary[14][41] = false
	summary[14][48] = true
	summary[53] = map[int]bool{}
	summary[53][60] = false

	if summary.GetBranchesTotal() != 4 {
		t.Fatalf("Expected summary branches total to be %d, but got %d", 4, summary.GetBranchesTotal())
	}
}

func TestBranchSummaryGetBranchesCovered(t *testing.T) {
	summary := statistic.BranchSummary{}
	summary[9] = map[int]bool{}
	summary[9][11] = true
	summary[14] = map[int]bool{}
	summary[14][41] = false
	summary[14][48] = true
	summary[53] = map[int]bool{}
	summary[53][60] = false

	if summary.GetBranchesCovered() != 2 {
		t.Fatalf("Expected summary branches total to be %d, but got %d", 2, summary.GetBranchesCovered())
	}
}
