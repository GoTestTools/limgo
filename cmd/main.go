package main

import (
	"fmt"
	"io"
	"os"

	"github.com/GoTestTools/limgo/pkg/domain"
	"github.com/GoTestTools/limgo/pkg/dto"
	"github.com/GoTestTools/limgo/pkg/flags"
)

func main() {
	cliFlags := flags.ParseFlags()

	covFile, err := os.Open(cliFlags.CoverageFile)
	if err != nil {
		fmt.Printf("Failed to open coverage file: %v\n", err)
		os.Exit(1)
	}
	defer covFile.Close()

	pf := domain.ParseFile(covFile)
	if len(pf) == 0 {
		fmt.Println("Empty test coverage file, aborting early")
		os.Exit(1)
	}

	confFile, err := os.Open(cliFlags.ConfigFile)
	if err != nil {
		fmt.Printf("Failed to open config: %v\n", err)
		os.Exit(1)
	}
	defer confFile.Close()

	cfg, err := dto.FromJSONString(confFile)
	if err != nil {
		fmt.Printf("Failed to parse config: %v\n", err)
		os.Exit(1)
	}

	module, err := domain.ParseModule(".", domain.AnalyzeModule)
	if err != nil {
		fmt.Printf("Failed to get go module: %v\n", err)
		os.Exit(1)
	}

	covStatistic := domain.BuildStatementStatistic(module, pf)

	printer := domain.NewPrinter(getOutput(cliFlags.OutputFile))
	printer.PrintStatistic(*covStatistic, cliFlags.Verbosity)

	covErrs, err := domain.Evaluate(*covStatistic, cfg)
	if err != nil {
		fmt.Printf("Failed to apply configured thresholds: %v\n", err)
	}

	if len(covErrs) != 0 {
		printer.PrintCoverageError(covErrs)
		os.Exit(1)
	}
	os.Exit(0)
}

func getOutput(outFile string) io.Writer {
	if outFile == "" {
		return os.Stdout
	}

	file, err := os.Create(outFile)
	if err != nil {
		fmt.Printf("Failed creating/opening output file: %v\n", err)
		return os.Stdout
	}
	return file
}
