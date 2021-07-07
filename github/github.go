package github

import (
	"regexp"
	"strings"
)

type GitHub struct{}

var re = regexp.MustCompile("^[-0-9A-Za-z]{3,39}$")

func containsNoIllegalPattern(username string) bool {
	return !strings.Contains(username, "--") &&
		!strings.HasPrefix(username, "-") &&
		!strings.HasSuffix(username, "-")
}

func looksGood(username string) bool {
	return re.MatchString(username)
}

func IsValid(username string) bool {
	return looksGood(username) && containsNoIllegalPattern(username)
}
