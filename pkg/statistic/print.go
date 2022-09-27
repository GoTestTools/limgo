package statistic

import (
	"fmt"
	"io"
	"text/tabwriter"
)

func PrintPretty(output io.Writer, stats *CoverageStatistic, verbosity uint) {
	if stats == nil || verbosity == 0 {
		return
	}

	w := tabwriter.NewWriter(output, 1, 1, 1, ' ', 0)
	printOverallCoverage(w, stats)

	if verbosity > 1 {
		printDirectoryCoverage(w, stats, verbosity)
	}

	w.Flush()
}

func printDirectoryCoverage(w *tabwriter.Writer, stats *CoverageStatistic, verbosity uint) {
	currentDir := ""
	for _, fileStat := range stats.FilesStatistics {
		if currentDir != fileStat.Dir {
			fmt.Fprintf(w, "  %s\t%s\t\n", fileStat.Dir, fmtPercentage(stats.GetCoveragePercent(fileStat.Dir)))
			currentDir = fileStat.Dir
		}

		if verbosity == 3 {
			printFileCoverageWithUncoveredLines(w, &fileStat)
		}
		if verbosity > 3 {
			printFileCoverage(w, &fileStat)
			printFuncCoverage(w, &fileStat)
		}
	}
}

func printFileCoverage(w *tabwriter.Writer, filestats *FileStatistic) {
	fmt.Fprintf(w, "    %s\t%v\t\n", filestats.Name, fmtPercentage(filestats.GetCoveragePercent()))
}
func printFileCoverageWithUncoveredLines(w *tabwriter.Writer, filestats *FileStatistic) {
	fmt.Fprintf(w, "    %s\t%v\t%v\n", filestats.Name, fmtPercentage(filestats.GetCoveragePercent()), filestats.GetUncoveredLines())
}

func printFuncCoverage(w *tabwriter.Writer, filestats *FileStatistic) {
	for _, fun := range filestats.FuncStats {
		fmt.Fprintf(w, "      %s\t%s\t%v\n", fun.Name, fmtPercentage(fun.GetCoveragePercent()), fun.GetUncoveredLines())
	}
}

func printOverallCoverage(w *tabwriter.Writer, stats *CoverageStatistic) {
	fmt.Fprintf(w, "Files\t%% Stmts\tUncovered Lines\n")
	fmt.Fprintf(w, "%s\t%s\t%v\n", stats.GlobalStatistic.ModuleName, fmtPercentage(stats.GlobalStatistic.GetCoveragePercent()), "")
}

func fmtPercentage(percentage float64) string {
	return fmt.Sprintf("%.2f%%", percentage)
}
