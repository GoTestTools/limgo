package main

import (
	"fmt"
	"io"
	"os"

	"github.com/GoTestTools/limgo/pkg"
	"github.com/GoTestTools/limgo/pkg/config"
	"github.com/GoTestTools/limgo/pkg/coverage"
	"github.com/GoTestTools/limgo/pkg/evaluation"
	"github.com/GoTestTools/limgo/pkg/flags"
	"github.com/GoTestTools/limgo/pkg/statistic"
)

func main() {
	cliFlags := flags.ParseFlags()
	if cliFlags.CoverageFile == "" {
		fmt.Printf("Missing argument -%s\n", flags.FlagCovFile)
		os.Exit(1)
	}

	covFile, err := os.Open(cliFlags.CoverageFile)
	if err != nil {
		fmt.Printf("Failed to open coverage file: %v\n", err)
		os.Exit(1)
	}
	defer covFile.Close()

	pf, err := coverage.Parse(covFile)
	if err != nil {
		fmt.Printf("Failed to parse coverage file: %v\n", err)
		os.Exit(1)
	}

	confFile, err := os.Open(cliFlags.ConfigFile)
	if err != nil {
		fmt.Printf("Failed to open config: %v\n", err)
		os.Exit(1)
	}
	defer confFile.Close()

	cfg, err := config.FromJSONString(confFile)
	if err != nil {
		fmt.Printf("Failed to parse config: %v\n", err)
		os.Exit(1)
	}

	module, err := pkg.GetCurrentModule()
	if err != nil {
		fmt.Printf("Failed to get go module: %v\n", err)
		os.Exit(1)
	}
	module.WithCoverage(pf).WithConfig(cfg)

	covStatistic, err := statistic.BuildCoverageStatistic(module)
	if err != nil {
		fmt.Printf("Failed to build coverage statistic: %v\n", err)
		os.Exit(1)
	}

	output := getOutput(cliFlags.OutputFile)
	covStatistic.Print(output, cliFlags.Verbosity)

	covErrs, err := evaluation.Evaluate(covStatistic, cfg)
	if err != nil {
		fmt.Printf("Failed to apply configured thresholds: %v\n", err)
	}

	if len(covErrs) != 0 {
		evaluation.PrintPretty(output, covErrs)
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
