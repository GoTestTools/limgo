package dto

import (
	"bytes"
	"fmt"
	"strings"
	"text/tabwriter"
)

type Formatter func(OutputDTO) (string, error)

func JSONFormat(output OutputDTO) (string, error) {
	var buf bytes.Buffer
	err := ToJSON(output, &buf)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func TabFormat(output OutputDTO) (string, error) {
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 1, 1, 1, ' ', 0)

	if output.Statistic != nil {

		fmt.Fprintf(w, "Files\t%% Stmt\t%% Line\t%% Branch\tUncovered Lines\n")
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%v\n",
			output.Statistic.Module,
			output.Statistic.Coverage.Stmts,
			output.Statistic.Coverage.Lines,
			output.Statistic.Coverage.Branches,
			"")
		for dirName, directory := range output.Statistic.Directories {
			fmt.Fprintf(w, "  %s\t%s\t%s\t%s\t%v\n",
				dirName,
				directory.Coverage.Stmts,
				directory.Coverage.Lines,
				directory.Coverage.Branches,
				"")
			for fileName, file := range directory.Files {
				fmt.Fprintf(w, "    %s\t%s\t%s\t%s\t%v\n",
					fileName,
					file.Coverage.Stmts,
					file.Coverage.Lines,
					file.Coverage.Branches,
					file.Coverage.Uncovered,
				)

				for fName, f := range file.Funcs {
					fmt.Fprintf(w, "      func %s\t%s\t%s\t%s\t%v\n",
						fName,
						f.Stmts,
						f.Lines,
						f.Branches,
						f.Uncovered,
					)
				}
			}
		}
	}

	if len(output.CoverageErrors) > 0 {
		fmt.Fprint(w, "\t\t\t\t\n")
		fmt.Fprint(w, colorizedFPrintln("\nExpected coverage thresholds not met:", ColorRed))
		for _, err := range output.CoverageErrors {
			txt := fmt.Sprintf(
				"   '%s':\texpected coverage threshold for %s of %s%%, but only got %s%%",
				err.AffectedFile, err.Type, err.ExpectedThreshold, err.ActualCovered,
			)
			fmt.Fprint(w, colorizedFPrintln(txt, ColorRed))
		}
	}

	//nolint:errcheck,gosec
	w.Flush()

	return buf.String(), nil
}

func MarkdownFormat(output OutputDTO) (string, error) {
	var buf strings.Builder

	if output.Statistic != nil {

		lineTmpl := "|%s|%s|%s|%s|%v|\n"
		buf.WriteString("| Files | % Stmt | % Line | % Branch | Uncovered Lines |\n")
		buf.WriteString("|-|-|-|-|-|\n")
		buf.WriteString(fmt.Sprintf(lineTmpl,
			output.Statistic.Module,
			output.Statistic.Coverage.Stmts,
			output.Statistic.Coverage.Lines,
			output.Statistic.Coverage.Branches,
			""))

		for dirName, directory := range output.Statistic.Directories {
			buf.WriteString(fmt.Sprintf(lineTmpl,
				"&nbsp;&nbsp;&nbsp;&nbsp;"+dirName,
				directory.Coverage.Stmts,
				directory.Coverage.Lines,
				directory.Coverage.Branches,
				""))
			for fileName, file := range directory.Files {
				buf.WriteString(fmt.Sprintf(lineTmpl,
					"&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"+fileName,
					file.Coverage.Stmts,
					file.Coverage.Lines,
					file.Coverage.Branches,
					file.Coverage.Uncovered),
				)

				for fName, f := range file.Funcs {
					buf.WriteString(fmt.Sprintf(lineTmpl,
						"&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;func "+fName,
						f.Stmts,
						f.Lines,
						f.Branches,
						f.Uncovered),
					)
				}
			}
		}
	}

	if len(output.CoverageErrors) > 0 {
		buf.WriteString("\n")
		buf.WriteString("\nExpected coverage thresholds not met:\n")
		for _, err := range output.CoverageErrors {
			txt := fmt.Sprintf(
				"- '%s':\texpected coverage threshold for %s of %s%%, but only got %s%%\n",
				err.AffectedFile, err.Type, err.ExpectedThreshold, err.ActualCovered,
			)
			buf.WriteString(txt)
		}
	}

	return buf.String(), nil
}

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

func colorizedFPrintln(txt string, color Color) string {
	return fmt.Sprintf("%s%s%s\n", color, txt, colorReset)
}
