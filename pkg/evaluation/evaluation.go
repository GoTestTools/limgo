package evaluation

import (
	"path"
	"regexp"

	"github.com/GoTestTools/limgo/pkg/config"
	"github.com/GoTestTools/limgo/pkg/statistic"
)

func Evaluate(cs *statistic.CoverageStatistic, cfg config.CoverageConfig) (errs []CoverageError, err error) {
	if cfg.Global != nil {

		if cfg.Global.Statements > cs.GlobalStatistic.GetCoveragePercent() {
			errs = append(errs, CoverageError{
				Type:              CoverageErrorStmt,
				AffectedFile:      cs.GlobalStatistic.ModuleName,
				ExpectedThreshold: cfg.Global.Statements,
				ActualCovered:     cs.GlobalStatistic.GetCoveragePercent(),
			})
		}

		if cfg.Global.Functions > cs.FilesStatistics.GetCoveragePercent() {
			errs = append(errs, CoverageError{
				Type:              CoverageErrorFunc,
				AffectedFile:      cs.GlobalStatistic.ModuleName,
				ExpectedThreshold: cfg.Global.Functions,
				ActualCovered:     cs.FilesStatistics.GetCoveragePercent(),
			})
		}
	}

	if cfg.Matcher != nil {

		regexMap := map[string]*regexp.Regexp{}
		for matcher := range *cfg.Matcher {
			r, err := regexp.Compile(matcher)
			if err != nil {
				return nil, err
			}
			regexMap[matcher] = r
		}

		for _, fileStat := range cs.FilesStatistics {
			for matcher, threshold := range *cfg.Matcher {
				fullPath := path.Join(fileStat.Dir, fileStat.Name)
				if regexMap[matcher].MatchString(fullPath) {
					if threshold.Statements > fileStat.GetCoveragePercent() {
						errs = append(errs, CoverageError{
							Type:              CoverageErrorStmt,
							AffectedFile:      fullPath,
							ExpectedThreshold: threshold.Statements,
							ActualCovered:     fileStat.GetCoveragePercent(),
						})
					}

					if threshold.Functions > fileStat.FuncStats.GetCoveragePercent() {
						errs = append(errs, CoverageError{
							Type:              CoverageErrorFunc,
							AffectedFile:      fullPath,
							ExpectedThreshold: threshold.Functions,
							ActualCovered:     fileStat.FuncStats.GetCoveragePercent(),
						})
					}
				}
			}
		}
	}

	return errs, nil
}
