package main

import (
	"fmt"
	"os"

	"github.com/engelmi/limgo/pkg"
	"github.com/engelmi/limgo/pkg/config"
	"github.com/engelmi/limgo/pkg/coverage"
	"github.com/engelmi/limgo/pkg/evaluation"
	"github.com/engelmi/limgo/pkg/flags"
	"github.com/engelmi/limgo/pkg/statistic"
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

	covStatistic.Print(os.Stdout, cliFlags.Verbosity)

	covErrs, err := evaluation.Evaluate(covStatistic, cfg)
	if err != nil {
		fmt.Printf("Failed to apply configured thresholds: %v\n", err)
	}

	if len(covErrs) != 0 {
		evaluation.PrintPretty(os.Stdout, covErrs)
		os.Exit(1)
	}
	os.Exit(0)
}