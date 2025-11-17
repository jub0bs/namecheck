package github

import (
	"fmt"
	"net/http"
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

func IsAvailable(username string) (bool, error) {
	resp, err := http.Get("https://github.com/" + username)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusOK:
		return false, nil
	case http.StatusNotFound:
		return true, nil
	default:
		return false, fmt.Errorf("github: couldn't check availability of %q", username)
	}
}
