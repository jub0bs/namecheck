package twitter

import (
	"regexp"
	"strings"
)

type Twitter struct{}

var re = regexp.MustCompile("^[0-9A-Z_a-z]{4,15}$")

func containsNoIllegalPattern(username string) bool {
	return !strings.Contains(strings.ToLower(username), "twitter")
}

func looksGood(username string) bool {
	return re.MatchString(username)
}

func (*Twitter) IsValid(username string) bool {
	return looksGood(username) && containsNoIllegalPattern(username)
}
