package domain

import (
	"fmt"
	"path"
	"regexp"

	"github.com/GoTestTools/limgo/pkg/dto"
	"github.com/GoTestTools/limgo/pkg/model/evaluation"
	"github.com/GoTestTools/limgo/pkg/model/statistic"
	"github.com/go-errors/errors"
)

func Evaluate(moduleStatistic statistic.ModuleStatistic, cfg dto.CoverageConfig) (errs []evaluation.CoverageError, err error) {
	if cfg.Global != nil {

		gblStmtCoverage := moduleStatistic.GetStmtCoverage()
		if cfg.Global.Statements > gblStmtCoverage {
			errs = append(errs, evaluation.CoverageError{
				Type:              evaluation.CoverageErrorStmt,
				AffectedFile:      moduleStatistic.Name,
				ExpectedThreshold: cfg.Global.Statements,
				ActualCovered:     gblStmtCoverage,
			})
		}

		gblLineCoverage := moduleStatistic.GetStmtCoverage()
		if cfg.Global.Lines > gblLineCoverage {
			errs = append(errs, evaluation.CoverageError{
				Type:              evaluation.CoverageErrorLines,
				AffectedFile:      moduleStatistic.Name,
				ExpectedThreshold: cfg.Global.Lines,
				ActualCovered:     gblLineCoverage,
			})
		}

		gblBranchesCoverage := moduleStatistic.GetBranchesCoverage()
		if cfg.Global.Lines > gblBranchesCoverage {
			errs = append(errs, evaluation.CoverageError{
				Type:              evaluation.CoverageErrorBranches,
				AffectedFile:      moduleStatistic.Name,
				ExpectedThreshold: cfg.Global.Branches,
				ActualCovered:     gblBranchesCoverage,
			})
		}

		// TODO: function coverage
	}

	if cfg.Matcher != nil {
		regexMap, err := compileMatcherRegex(cfg)
		if err != nil {
			return nil, err
		}

		for matcher, threshold := range *cfg.Matcher {
			filteredFiles := moduleStatistic.FilterFileStatistics(regexMap[matcher])

			for _, filteredFile := range filteredFiles {
				filePath := path.Join(filteredFile.Directory, filteredFile.Name)

				stmtCoverage := filteredFile.GetStmtCoverage()
				if threshold.Statements > stmtCoverage {
					errs = append(errs, evaluation.CoverageError{
						Type:              evaluation.CoverageErrorStmt,
						AffectedFile:      filePath,
						ExpectedThreshold: threshold.Statements,
						ActualCovered:     stmtCoverage,
					})
				}

				gblLineCoverage := filteredFile.GetStmtCoverage()
				if threshold.Lines > gblLineCoverage {
					errs = append(errs, evaluation.CoverageError{
						Type:              evaluation.CoverageErrorLines,
						AffectedFile:      filePath,
						ExpectedThreshold: threshold.Lines,
						ActualCovered:     gblLineCoverage,
					})
				}

				gblBranchesCoverage := filteredFile.GetBranchesCoverage()
				if threshold.Lines > gblBranchesCoverage {
					errs = append(errs, evaluation.CoverageError{
						Type:              evaluation.CoverageErrorBranches,
						AffectedFile:      filePath,
						ExpectedThreshold: threshold.Branches,
						ActualCovered:     gblBranchesCoverage,
					})
				}
			}
		}
	}

	return errs, nil
}

func compileMatcherRegex(cfg dto.CoverageConfig) (map[string]*regexp.Regexp, error) {
	regexMap := map[string]*regexp.Regexp{}
	for matcher := range *cfg.Matcher {
		r, err := regexp.Compile(matcher)
		if err != nil {
			return nil, errors.New(fmt.Errorf("failed to compile regex '%s': %w", matcher, err))
		}
		regexMap[matcher] = r
	}
	return regexMap, nil
}
