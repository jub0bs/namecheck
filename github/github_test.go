package github_test // external test package: black-box testing

import (
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
		"starts with a hyphen":             {"-jub0bs", false},
		"ends with a hyphen":               {"jub0bs-", false},
		"contains illegal chars":           {"jub&obs", false},
		"all good":                         {"jub0bs", true},
		// other test cases...
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

func (sc *StubClient) Get(_ string) (*http.Response, error) {
	if sc.Err != nil {
		return nil, sc.Err
	}
	res := http.Response{
		StatusCode: sc.StatusCode,
		Body:       http.NoBody,
	}
	return &res, nil
}

func TestIsAvailable404(t *testing.T) {
	gh := github.GitHub{
		Client: &StubClient{StatusCode: http.StatusNotFound},
	}
	avail, err := gh.IsAvailable("whatever")
	if err != nil || !avail {
		t.Errorf("got %t, %v; want true, nil", avail, err)
	}
}
