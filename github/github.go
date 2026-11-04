package github

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/jub0bs/namecheck"
)

var re = regexp.MustCompile("^[A-Za-z0-9-]{3,39}$")

type GitHub struct {
	Client namecheck.Doer
}

func (*GitHub) IsValid(username string) bool {
	return !strings.HasPrefix(username, "-") &&
		!strings.HasSuffix(username, "-") &&
		!strings.Contains(username, "--") &&
		re.MatchString(username)
}

func (gh *GitHub) IsAvailable(username string) (bool, error) {
	u, err := url.JoinPath("https://github.com", username)
	if err != nil {
		return false, err
	}
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return false, err
	}
	resp, err := gh.Client.Do(req)
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
		return false, fmt.Errorf("unexpected status %d", resp.StatusCode)
	}
}

func (*GitHub) String() string {
	return "GitHub"
}
