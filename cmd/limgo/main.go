package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/GoTestTools/limgo/cmd/flags"
	"github.com/GoTestTools/limgo/pkg/domain"
	"github.com/GoTestTools/limgo/pkg/dto"
)

func main() {
	cliFlags := flags.ParseFlags()

	covFile, err := os.Open(cliFlags.CoverageFile)
	if err != nil {
		fmt.Printf("Failed to open coverage file: %v\n", err)
		os.Exit(1)
	}
	//nolint:errcheck,gosec
	defer covFile.Close()

	pf := domain.ParseFile(covFile)
	if len(pf) == 0 {
		fmt.Println("Empty test coverage file, aborting early")
		//nolint:gocritic
		os.Exit(1)
	}

	confFile, err := os.Open(cliFlags.ConfigFile)
	if err != nil {
		fmt.Printf("Failed to open config: %v\n", err)
		os.Exit(1)
	}
	//nolint:errcheck,gosec
	defer confFile.Close()

	cfg, err := dto.ConfigFromJSONString(confFile)
	if err != nil {
		fmt.Printf("Failed to parse config: %v\n", err)
		os.Exit(1)
	}

	module, err := domain.ParseModule(".", domain.NewModuleAnalyzer(cfg.StatisticConfig.Excludes))
	if err != nil {
		fmt.Printf("Failed to get go module: %v\n", err)
		os.Exit(1)
	}

	covStatistic := domain.BuildStatementStatistic(module, pf)
	covErrs, err := domain.Evaluate(*covStatistic, cfg.CoverageConfig)
	if err != nil {
		fmt.Printf("Failed to apply configured thresholds: %v\n", err)
	}

	printer := domain.NewPrinter(getOutput(cliFlags.OutputFile), getFormatter(cliFlags.OutputFormat))
	printer.Print(covStatistic, covErrs, cliFlags.Verbosity)

	if len(covErrs) != 0 {
		os.Exit(1)
	}
	os.Exit(0)
}

func getOutput(outFile string) io.Writer {
	if outFile == "" {
		return os.Stdout
	}

	file, err := os.Create(filepath.Clean(outFile))
	if err != nil {
		fmt.Printf("Failed creating/opening output file: %v\n", err)
		return os.Stdout
	}
	return file
}

func getFormatter(outFormat string) dto.Formatter {
	switch outFormat {
	case "tab":
		return dto.TabFormat
	case "json":
		return dto.JSONFormat
	case "md":
		return dto.MarkdownFormat
	default:
		fmt.Printf("No formatter found for '%s', falling back to '%s'", outFormat, "tab")
		return dto.TabFormat
	}
}
