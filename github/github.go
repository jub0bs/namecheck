package github

import (
	"errors"
	"net/http"
	"regexp"
	"strings"
)

var re = regexp.MustCompile("^[A-Z0-9a-z-]{3,39}$")

type GitHub struct {
	Client *http.Client
}

func (*GitHub) IsValid(username string) bool {
	return !strings.HasPrefix(username, "-") &&
		!strings.HasSuffix(username, "-") &&
		!strings.Contains(username, "--") &&
		re.MatchString(username)
}

func (gh *GitHub) IsAvailable(username string) (bool, error) {
	addr := "https://github.com/" + username
	resp, err := gh.Client.Get(addr)
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

func (gh *GitHub) String() string {
	return "GitHub"
}
