package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-errors/errors"
)

func ToJSON(obj interface{}, w io.Writer) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")

	err := encoder.Encode(obj)
	if err != nil {
		return errors.New(fmt.Errorf("failed to marshal to json: %w", err))
	}
	return nil
}

func FromJSONString(obj interface{}, r io.Reader) error {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&obj)
	if err != nil {
		return errors.New(fmt.Errorf("failed to unmarshal from string: %w", err))
	}

	return nil
}
