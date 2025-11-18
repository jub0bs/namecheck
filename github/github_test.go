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
		"too short":                        {"ab", false},
		"too long":                         {strings.Repeat("a", 40), false},
		"contains illegal chars":           {"^^^", false},
		"starts with a hyphen":             {"-jub0bs", false},
		"ends with a hyphen":               {"jub0bs-", false},
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

func (c *StubClient) Get(url string) (*http.Response, error) {
	if c.Err != nil {
		return nil, c.Err
	}
	resp := http.Response{
		StatusCode: c.StatusCode,
		Body:       http.NoBody,
	}
	return &resp, nil
}

func TestIsAvailableError(t *testing.T) {
	gh := github.GitHub{
		&StubClient{Err: errors.New("oh no")},
	}
	got, err := gh.IsAvailable("whatever")
	if err == nil || got {
		t.Errorf("IsAvailable(...): got %t, %v; want false, some non-nil error", got, err)
	}
}

func TestIsAvailable404(t *testing.T) {
	gh := github.GitHub{
		&StubClient{StatusCode: http.StatusNotFound},
	}
	got, err := gh.IsAvailable("whatever")
	if err != nil || !got {
		t.Errorf("IsAvailable(...): got %t, %v; want true, nil", got, err)
	}
}
