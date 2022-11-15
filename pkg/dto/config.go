package dto

import "io"

type LimgoConfig struct {
	CoverageConfig  `json:"coverage"`
	StatisticConfig `json:"statistic"`
}

type CoverageConfig struct {
	Global   *Threshold `json:"global,omitempty"`
	Matcher  *Matcher   `json:"matcher,omitempty"`
	Excludes []string   `json:"excludes"`
}

type Matcher map[string]Threshold

type Threshold struct {
	Statements float64 `json:"statements"`
	Lines      float64 `json:"lines"`
	Branches   float64 `json:"branches"`
	Functions  float64 `json:"functions"`
}

type StatisticConfig struct {
	Excludes []string `json:"excludes"`
}

func (config LimgoConfig) ToJSON(w io.Writer) error {
	setConfigDefaults(&config)
	return ToJSON(config, w)
}

func ConfigFromJSONString(r io.Reader) (LimgoConfig, error) {
	var config LimgoConfig
	err := FromJSONString(&config, r)
	if err != nil {
		return LimgoConfig{}, err
	}

	setConfigDefaults(&config)
	return config, nil
}

func setConfigDefaults(config *LimgoConfig) {
	// exclude vendor directory by default for coverage and statistic
	vendorExcludePattern := "vendor/.*"
	isCoverageVendorExcluded := false
	for _, exclude := range config.CoverageConfig.Excludes {
		if exclude == vendorExcludePattern {
			isCoverageVendorExcluded = true
			break
		}
	}
	if !isCoverageVendorExcluded {
		config.CoverageConfig.Excludes = append(config.CoverageConfig.Excludes, vendorExcludePattern)
	}

	isStatisticVendorExcluded := false
	for _, exclude := range config.StatisticConfig.Excludes {
		if exclude == vendorExcludePattern {
			isStatisticVendorExcluded = true
			break
		}
	}
	if !isStatisticVendorExcluded {
		config.StatisticConfig.Excludes = append(config.StatisticConfig.Excludes, vendorExcludePattern)
	}
}
