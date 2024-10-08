package github

import (
	"errors"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

var re = regexp.MustCompile("^[a-zA-Z0-9-]{3,39}$")

func IsValid(username string) bool {
	return !strings.HasPrefix(username, "-") &&
		!strings.HasSuffix(username, "-") &&
		!strings.Contains(username, "--") &&
		re.MatchString(username)
}

func IsAvailable(username string) (bool, error) {
	addr, err := url.JoinPath("https://github.com", url.PathEscape(username))
	if err != nil {
		return false, err
	}
	resp, err := http.Get(addr)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusNotFound:
		return true, nil
	case http.StatusOK:
		return false, nil
	default:
		return false, errors.New("unknown availability")
	}
}
