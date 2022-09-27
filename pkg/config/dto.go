package config

type CoverageConfig struct {
	CoverageThreshold `json:"coverageThreshold"`
	Excludes          []string `json:"excludes"`
}

type CoverageThreshold struct {
	Global  *Threshold `json:"global,omitempty"`
	Matcher *Matcher   `json:"matcher,omitempty"`
}

type Matcher map[string]Threshold

type Threshold struct {
	Statements float64 `json:"statements"`
	Functions  float64 `json:"functions"`
	//Branches   float64 `json:"branches"`
	//Lines      float64 `json:"lines"`
}
