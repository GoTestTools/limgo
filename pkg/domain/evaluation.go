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
	errs = append(errs, evaluateGlobal(moduleStatistic, cfg.Global)...)

	matcherErrs, err := evaluateMatcher(moduleStatistic, cfg.Matcher)
	if err != nil {
		return nil, err
	}
	errs = append(errs, matcherErrs...)

	return errs, nil
}

func evaluateGlobal(moduleStatistic statistic.ModuleStatistic, glblCfg *dto.Threshold) (errs []evaluation.CoverageError) {
	if glblCfg != nil {

		gblStmtCoverage := moduleStatistic.GetStmtCoverage()
		if glblCfg.Statements > gblStmtCoverage {
			errs = append(errs, evaluation.CoverageError{
				Type:              evaluation.CoverageErrorStmt,
				AffectedFile:      moduleStatistic.Name,
				ExpectedThreshold: glblCfg.Statements,
				ActualCovered:     gblStmtCoverage,
			})
		}

		gblLineCoverage := moduleStatistic.GetStmtCoverage()
		if glblCfg.Lines > gblLineCoverage {
			errs = append(errs, evaluation.CoverageError{
				Type:              evaluation.CoverageErrorLines,
				AffectedFile:      moduleStatistic.Name,
				ExpectedThreshold: glblCfg.Lines,
				ActualCovered:     gblLineCoverage,
			})
		}

		gblBranchesCoverage := moduleStatistic.GetBranchesCoverage()
		if glblCfg.Lines > gblBranchesCoverage {
			errs = append(errs, evaluation.CoverageError{
				Type:              evaluation.CoverageErrorBranches,
				AffectedFile:      moduleStatistic.Name,
				ExpectedThreshold: glblCfg.Branches,
				ActualCovered:     gblBranchesCoverage,
			})
		}
	}

	return errs
}

func evaluateMatcher(moduleStatistic statistic.ModuleStatistic, matcherCfg *dto.Matcher) (errs []evaluation.CoverageError, err error) {
	if matcherCfg != nil {
		regexMap, err := compileMatcherRegex(matcherCfg)
		if err != nil {
			return nil, err
		}

		for matcher, threshold := range *matcherCfg {
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

func compileMatcherRegex(matcherCfg *dto.Matcher) (map[string]*regexp.Regexp, error) {
	regexMap := map[string]*regexp.Regexp{}
	for matcher := range *matcherCfg {
		r, err := regexp.Compile(matcher)
		if err != nil {
			return nil, errors.New(fmt.Errorf("failed to compile regex '%s': %w", matcher, err))
		}
		regexMap[matcher] = r
	}
	return regexMap, nil
}
