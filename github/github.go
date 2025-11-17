package github

import (
	"regexp"
	"strings"
)

func IsValid(username string) bool {
	return !strings.HasPrefix(username, "-") &&
		!strings.HasSuffix(username, "-") &&
		!strings.Contains(username, "--") &&
		re.MatchString(username)
}

var re = regexp.MustCompile("^[a-zA-Z0-9-]{3,39}$")
