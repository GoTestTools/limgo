package domain

import (
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/GoTestTools/limgo/pkg/model/evaluation"
	"github.com/GoTestTools/limgo/pkg/model/statistic"
)

type Color string

// ASCII escape codes for colors, see: https://en.wikipedia.org/wiki/ANSI_escape_code#Colors
const (
	colorReset Color = "\033[0m"

	ColorRed    Color = "\033[31m"
	ColorGreen  Color = "\033[32m"
	ColorYellow Color = "\033[33m"
	ColorBlue   Color = "\033[34m"
	ColorPurple Color = "\033[35m"
	ColorCyan   Color = "\033[36m"
	ColorWhite  Color = "\033[37m"
)

type printer struct {
	w io.Writer
}

func NewPrinter(w io.Writer) printer {
	return printer{w: w}
}

func (p printer) PrintCoverageError(errs []evaluation.CoverageError) {
	w := tabwriter.NewWriter(p.w, 1, 1, 1, ' ', 0)

	p.colorizedFPrintln(w, "\nExpected coverage thresholds not met:", ColorRed)
	for _, err := range errs {
		txt := fmt.Sprintf(
			"   '%s':\texpected coverage threshold for %s of %.2f%%, but only got %.2f%%",
			err.AffectedFile, err.Type, err.ExpectedThreshold, err.ActualCovered,
		)
		p.colorizedFPrintln(w, txt, ColorRed)

	}

	w.Flush()
}

func (p printer) PrintStatistic(stats statistic.ModuleStatistic, verbosity uint) {
	if verbosity == 0 {
		return
	}

	w := tabwriter.NewWriter(p.w, 1, 1, 1, ' ', 0)
	p.printModuleCoverage(w, stats)

	if verbosity > 1 {
		p.printDirectoryCoverage(w, stats.GetDirectoryStatistics(), verbosity)
	}

	w.Flush()
}

func (p printer) printModuleCoverage(w *tabwriter.Writer, stats statistic.ModuleStatistic) {
	fmt.Fprintf(w, "Files\t%% Stmt\t%% Line\t%% Branch\tUncovered Lines\n")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%v\n",
		stats.Name,
		fmtPercentage(stats.GetStmtCoverage()),
		fmtPercentage(stats.GetLinesCoverage()),
		fmtPercentage(stats.GetBranchesCoverage()),
		"")
}

func (p printer) printDirectoryCoverage(w *tabwriter.Writer, stats []statistic.DirectoryStatistic, verbosity uint) {
	for _, directoryStat := range stats {
		fmt.Fprintf(w, "  %s\t%s\t%s\t%s\t%v\n",
			directoryStat.GetName(),
			fmtPercentage(directoryStat.GetStmtCoverage()),
			fmtPercentage(directoryStat.GetLinesCoverage()),
			fmtPercentage(directoryStat.GetBranchesCoverage()),
			"")

		if verbosity > 2 {
			p.printFileCoverage(w, directoryStat.GetFileStatistics(), verbosity)
		}
	}
}

func (p printer) printFileCoverage(w *tabwriter.Writer, stats []statistic.FileStatistic, verbosity uint) {
	for _, fileStat := range stats {
		fmt.Fprintf(w, "    %s\t%s\t%s\t%s\t%v\n",
			fileStat.Name,
			fmtPercentage(fileStat.GetStmtCoverage()),
			fmtPercentage(fileStat.GetLinesCoverage()),
			fmtPercentage(fileStat.GetBranchesCoverage()),
			"")

		if verbosity > 3 {
			p.printFuncCoverage(w, fileStat.FunctionStatistics, verbosity)
		}
	}
}

func (p printer) printFuncCoverage(w *tabwriter.Writer, stats []statistic.FunctionStatistic, verbosity uint) {
	for _, funcStat := range stats {
		if funcStat.Name == "GenDecl" {
			continue
		}
		fmt.Fprintf(w, "      func %s\t%s\t%s\t%s\t%v\n",
			funcStat.Name,
			fmtPercentage(funcStat.GetStmtCoverage()),
			fmtPercentage(funcStat.GetLinesCoverage()),
			fmtPercentage(funcStat.GetBranchesCoverage()),
			"")
	}
}

func (p printer) colorizedFPrintln(w io.Writer, txt string, color Color) {
	fmt.Fprintf(w, "%s%s%s\n", color, txt, colorReset)
}

func fmtPercentage(p float64) string {
	return fmt.Sprintf("%.2f", p)
}
