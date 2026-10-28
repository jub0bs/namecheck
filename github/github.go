package github

import (
	"regexp"
	"strings"
)

var re = regexp.MustCompile("^[A-Za-z0-9-]{3,39}$")

func IsValid(username string) bool {
	return !strings.HasPrefix(username, "-") &&
		!strings.HasSuffix(username, "-") &&
		!strings.Contains(username, "--") &&
		re.MatchString(username)
}
