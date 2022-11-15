package dto_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/GoTestTools/limgo/pkg/dto"
	"github.com/google/go-cmp/cmp"
)

func TestConfigToJSON_Succeeds(t *testing.T) {
	cfg := dto.LimgoConfig{
		CoverageConfig: dto.CoverageConfig{
			Global: &dto.Threshold{
				Statements: 100,
				Functions:  100,
			},
			Matcher: &dto.Matcher{
				"pkg/coverage/line.go": dto.Threshold{
					Statements: 10,
					Functions:  20,
				},
			},
			Excludes: []string{"vendor/.*"},
		},
		StatisticConfig: dto.StatisticConfig{
			Excludes: []string{"vendor/.*"},
		},
	}

	cfgBuffer := &bytes.Buffer{}
	err := cfg.ToJSON(cfgBuffer)
	if err != nil {
		t.Fatalf("Unexpected error writing config struct to string: %v", err)
	}

	expectedRawCfg := `{
	"coverage": {
		"global": {
			"statements": 100,
			"lines": 0,
			"branches": 0,
			"functions": 100
		},
		"matcher": {
			"pkg/coverage/line.go": {
				"statements": 10,
				"lines": 0,
				"branches": 0,
				"functions": 20
			}
		},
		"excludes": [
			"vendor/.*"
		]
	},
	"statistic": {
		"excludes": [
			"vendor/.*"
		]
	}
}
`

	if diff := cmp.Diff(expectedRawCfg, cfgBuffer.String()); diff != "" {
		t.Fatalf("Detected difference in parsed config: %s", diff)
	}
}

func TestConfigToJSON_SucceedsSetsDefaults(t *testing.T) {
	cfg := dto.LimgoConfig{
		CoverageConfig:  dto.CoverageConfig{},
		StatisticConfig: dto.StatisticConfig{},
	}

	cfgBuffer := &bytes.Buffer{}
	err := cfg.ToJSON(cfgBuffer)
	if err != nil {
		t.Fatalf("Unexpected error writing config struct to string: %v", err)
	}

	expectedRawCfg := `{
	"coverage": {
		"excludes": [
			"vendor/.*"
		]
	},
	"statistic": {
		"excludes": [
			"vendor/.*"
		]
	}
}
`

	if diff := cmp.Diff(expectedRawCfg, cfgBuffer.String()); diff != "" {
		t.Fatalf("Detected difference in parsed config: %s", diff)
	}
}

func TestConfigFromJSONString_Succeeds(t *testing.T) {
	rawCfg := `{
	"coverage": {
		"global": {
			"statements": 100,
			"lines": 0,
			"branches": 0,
			"functions": 100
		},
		"matcher": {
			"pkg/coverage/line.go": {
				"statements": 10,
				"lines": 0,
				"branches": 0,
				"functions": 20
			}
		},
		"excludes": [
			"vendor/.*"
		]
	},
	"statistic": {
		"excludes": [
			"vendor/.*"
		]
	}
}
`

	parsedCfg, err := dto.ConfigFromJSONString(strings.NewReader(rawCfg))
	if err != nil {
		t.Fatalf("Unexpected error parsing config string to struct: %v", err)
	}

	expectedCfg := dto.LimgoConfig{
		CoverageConfig: dto.CoverageConfig{
			Global: &dto.Threshold{
				Statements: 100,
				Functions:  100,
			},
			Matcher: &dto.Matcher{
				"pkg/coverage/line.go": dto.Threshold{
					Statements: 10,
					Functions:  20,
				},
			},
			Excludes: []string{"vendor/.*"},
		},
		StatisticConfig: dto.StatisticConfig{
			Excludes: []string{"vendor/.*"},
		},
	}

	if diff := cmp.Diff(expectedCfg, parsedCfg); diff != "" {
		t.Fatalf("Detected difference in parsed config: %s", diff)
	}
}

func TestConfigFromJSONString_SucceedsSetsDefaults(t *testing.T) {
	rawCfg := `{}`

	parsedCfg, err := dto.ConfigFromJSONString(strings.NewReader(rawCfg))
	if err != nil {
		t.Fatalf("Unexpected error parsing config string to struct: %v", err)
	}

	expectedCfg := dto.LimgoConfig{
		CoverageConfig: dto.CoverageConfig{
			Excludes: []string{"vendor/.*"},
		},
		StatisticConfig: dto.StatisticConfig{
			Excludes: []string{"vendor/.*"},
		},
	}

	if diff := cmp.Diff(expectedCfg, parsedCfg); diff != "" {
		t.Fatalf("Detected difference in parsed config: %s", diff)
	}
}

func TestFromJSONString_FailsDueToInvalidJSON(t *testing.T) {
	cfgMissingBracket := `
	{
		"coverageThreshold": {
			"global": {
				"statements": 100,
				"functions": 100
			},
		}
	`

	parsedCfg, err := dto.ConfigFromJSONString(strings.NewReader(cfgMissingBracket))
	if err == nil {
		t.Fatalf("Expected error when parsing invalid config string, but got cfg: %v", parsedCfg)
	}
}
