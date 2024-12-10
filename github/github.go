package github

import (
	"errors"
	"net/http"
	"regexp"
	"strings"
)

var re = regexp.MustCompile("^[a-zA-Z0-9-]{3,39}$")

type GitHub struct {
	Client *http.Client
}

func (*GitHub) IsValid(username string) bool {
	return !strings.Contains(username, "--") &&
		!strings.HasPrefix(username, "-") &&
		!strings.HasSuffix(username, "-") &&
		re.MatchString(username)
}

func (gh *GitHub) IsAvailable(username string) (bool, error) {
	res, err := gh.Client.Get("https://github.com/" + username)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()
	switch res.StatusCode {
	case http.StatusNotFound:
		return true, nil
	case http.StatusOK:
		return false, nil
	default:
		return false, errors.New("github: unexpected status")
	}
}
