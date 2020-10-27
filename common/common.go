package common

import (
	"fmt"
	"regexp"
)

func Validate(input string) error {
	var invalid = regexp.MustCompile(`(\(.+\)+)|(.*;\s*(and|or)\s+.+)`)

	if invalid.Match([]byte(input)) {
		return fmt.Errorf("input string cannot initialize subshells or run commands: %v", input)
	}

	return nil
}
