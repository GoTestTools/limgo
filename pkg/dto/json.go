package dto

import (
	"encoding/json"
	"io"
)

func (config CoverageConfig) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	err := encoder.Encode(config)
	if err != nil {
		return err
	}
	return nil
}

func FromJSONString(r io.Reader) (CoverageConfig, error) {
	var config CoverageConfig
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&config)
	if err != nil {
		return CoverageConfig{}, err
	}
	return config, nil
}
