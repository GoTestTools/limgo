package dto

import (
	"encoding/json"
	"io"
)

func (config LimgoConfig) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")

	setDefaults(&config)
	err := encoder.Encode(config)
	if err != nil {
		return err
	}
	return nil
}

func FromJSONString(r io.Reader) (LimgoConfig, error) {
	var config LimgoConfig
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&config)
	if err != nil {
		return LimgoConfig{}, err
	}

	setDefaults(&config)
	return config, nil
}

func setDefaults(config *LimgoConfig) {
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
