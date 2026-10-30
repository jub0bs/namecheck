package github_test

import (
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/jub0bs/namecheck/github"
)

func TestIsValid(t *testing.T) {
	type TestCase struct {
		username string
		want     bool
	}
	testCases := map[string]TestCase{
		"contains two consecutive hyphens": {"jub0bs--on-GitHub", false},
		"starts with a hyphen":             {"-jub0bs", false},
		"ends with a hyphen":               {"jub0bs-", false},
		"contains invalid chars":           {"jub^bs", false},
		"too short":                        {"ju", false},
		"too long":                         {strings.Repeat("a", 40), false},
		"all good":                         {"jub0bs", true},
	}
	var gh github.GitHub
	for desc, tc := range testCases {
		f := func(t *testing.T) {
			got := gh.IsValid(tc.username)
			if got != tc.want {
				const tmpl = "github.IsValid(%q): got %t; want %t"
				t.Errorf(tmpl, tc.username, got, tc.want)
			}
		}
		t.Run(desc, f)
	}
}

type StubClient struct {
	StatusCode int
	Err        error
}

func (sc *StubClient) Do(req *http.Request) (*http.Response, error) {
	if sc.Err != nil {
		return nil, sc.Err
	}
	resp := http.Response{
		StatusCode: sc.StatusCode,
		Body:       http.NoBody,
	}
	return &resp, nil
}

func TestIsAvailableError(t *testing.T) {
	client := StubClient{
		Err: errors.New("oh no"),
	}
	gh := github.GitHub{
		Client: &client,
	}
	const username = "whatever"
	avail, err := gh.IsAvailable(username)
	if err == nil || avail {
		t.Errorf("IsAvailable(%q): got %t, %v; want false, some non-nil error", username, avail, err)
	}
}

func TestIsAvailableNotFound(t *testing.T) {
	client := StubClient{
		StatusCode: http.StatusNotFound,
	}
	gh := github.GitHub{
		Client: &client,
	}
	const username = "whatever"
	avail, err := gh.IsAvailable(username)
	if err != nil || !avail {
		t.Errorf("IsAvailable(%q): got %t, %v; want true, nil", username, avail, err)
	}
}
