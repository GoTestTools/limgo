package dto

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
