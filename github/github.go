package github

import (
	"regexp"
	"strings"
)

var re = regexp.MustCompile("^[A-Z0-9a-z-]{3,39}$")

func IsValid(username string) bool {
	return !strings.HasPrefix(username, "-") &&
		!strings.HasSuffix(username, "-") &&
		!strings.Contains(username, "--") &&
		re.MatchString(username)
}
