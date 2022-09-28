package coverage

import (
	"regexp"

	"github.com/go-errors/errors"
)

type LineMatcher struct {
	expr *regexp.Regexp
}

func NewLineMatcher() LineMatcher {
	return LineMatcher{
		expr: regexp.MustCompile(`(.+[^\/])\/(.+\.go):(\d+).(\d+),(\d+).(\d+)\s(\d+)\s(\d+)`),
	}
}

func (reg LineMatcher) Apply(line string) (ParsedLine, error) {
	if !reg.expr.MatchString(line) {
		return nil, errors.New("string does not match coverfile regex")
	}

	matches := reg.expr.FindStringSubmatch(line)
	if len(matches) != 9 {
		return nil, errors.Errorf("string did not yield the required number of matches (expected %d, got %d)", 9, len(matches))
	}

	// slice off match for whole string
	return matches[1:], nil
}
