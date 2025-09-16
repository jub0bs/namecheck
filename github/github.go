package github

import (
	"context"
	"errors"
	"net/http"
	"regexp"
	"strings"

	"github.com/jub0bs/namecheck"
)

var re = regexp.MustCompile("^[A-Z0-9a-z-]{3,39}$")

type GitHub struct {
	Client namecheck.Getter
}

func (*GitHub) IsValid(username string) bool {
	return !strings.HasPrefix(username, "-") &&
		!strings.HasSuffix(username, "-") &&
		!strings.Contains(username, "--") &&
		re.MatchString(username)
}

func (gh *GitHub) IsAvailable(ctx context.Context, username string) (bool, error) {
	addr := "https://github.com/" + username
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, addr, nil)
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
		return false, errors.New("unknown availability")
	}
}

func (gh *GitHub) String() string {
	return "GitHub"
}
