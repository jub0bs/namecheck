package github

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/jub0bs/namecheck"
)

type GitHub struct {
	Client namecheck.Doer
}

func (*GitHub) IsValid(username string) bool {
	return !strings.HasPrefix(username, "-") &&
		!strings.HasSuffix(username, "-") &&
		!strings.Contains(username, "--") &&
		re.MatchString(username)
}

var re = regexp.MustCompile("^[a-zA-Z0-9-]{3,39}$")

func (gh *GitHub) IsAvailable(username string) (bool, error) {
	addr := "https://github.com/" + username
	req, err := http.NewRequest(http.MethodGet, addr, nil)
	if err != nil {
		return false, err
	}
	resp, err := gh.Client.Do(req)
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

func (gh *GitHub) String() string {
	return "GitHub"
}
