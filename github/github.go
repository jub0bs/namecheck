package github

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/jub0bs/namecheck"
)

type GitHub struct {
	Client namecheck.Getter
}

var re = regexp.MustCompile("^[a-zA-Z0-9-]{3,39}$")

func (*GitHub) IsValid(username string) bool {
	if strings.HasPrefix(username, "-") ||
		strings.HasSuffix(username, "-") ||
		strings.Contains(username, "--") {
		return false
	}
	return re.MatchString(username)
}

func (gh *GitHub) IsAvailable(username string) (bool, error) {
	url := fmt.Sprintf("https://github.com/%s", username)
	resp, err := gh.Client.Get(url)
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
		return false, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
}

func (*GitHub) String() string {
	return "GitHub"
}
