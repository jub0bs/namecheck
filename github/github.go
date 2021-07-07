package github

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/jub0bs/namecheck"
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

func (*GitHub) IsValid(username string) bool {
	return looksGood(username) && containsNoIllegalPattern(username)
}

func (gh *GitHub) IsAvailable(username string) (bool, error) {
	endpoint := "https://github.com/" + username
	resp, err := http.Get(endpoint)
	if err != nil {
		err1 := namecheck.ErrUnknownAvailability{
			Username: username,
			Platform: gh.String(),
			Cause:    err,
		}
		return false, &err1
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusNotFound, nil
}

func (*GitHub) String() string {
	return "GitHub"
}
