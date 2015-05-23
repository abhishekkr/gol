package golbin

import (
	"regexp"
	"strings"
)

/*
Grep tries matching lines in a file for given pattern and return matched lines.
*/
func Grep(re, lines string) (string, error) {
	match := ""
	regex, regexErr := regexp.Compile(re)
	if regexErr != nil {
		return "", regexErr
	}

	for _, line := range strings.Split(lines, "\n") {
		if regex.MatchString(line) {
			match = match + "\n" + line
		}
	}
	return match, nil
}
