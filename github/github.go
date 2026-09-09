package github

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

var re = regexp.MustCompile("^[A-Za-z0-9-]{3,39}$")

func IsValid(username string) bool {
	return !strings.HasPrefix(username, "-") &&
		!strings.Contains(username, "--") &&
		!strings.HasSuffix(username, "-") &&
		re.MatchString(username)
}

func IsAvailable(username string) (bool, error) {
	req, err := http.NewRequest(http.MethodGet, "https://github.com/"+username, nil)
	if err != nil {
		return false, err
	}
	resp, err := http.DefaultClient.Do(req)
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
		return false, fmt.Errorf("unknown availability of %q on GitHub", username)
	}
}
