package github_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/jub0bs/namecheck/github"
	"github.com/jub0bs/namecheck/stub"
)

func TestIsValid(t *testing.T) {
	type TestCase struct {
		username string
		want     bool
	}
	testCases := map[string]TestCase{
		"contains two consecutive hyphens": {"jub0bs--on-GitHub", false},
		"contains illegal characters":      {"jub0bs on GitHub", false},
		"too short":                        {"ju", false},
		"too long":                         {strings.Repeat("a", 40), false},
		"all good":                         {"jub0bs", true},
		// other test cases...
	}
	var gh github.GitHub
	for desc, tc := range testCases {
		f := func(t *testing.T) {
			got := gh.IsValid(tc.username)
			if got != tc.want {
				t.Errorf("github.IsValid(%q): got %t; want %t", tc.username, got, tc.want)
			}
		}
		t.Run(desc, f)
	}
}

func TestIsAvailableNotFound(t *testing.T) {
	gh := github.GitHub{
		Client: &stub.SuccessfulGetter{StatusCode: http.StatusNotFound},
	}
	const username = "whatever"
	avail, err := gh.IsAvailable(username)
	if err != nil || !avail {
		const tmpl = "IsAvailable(%q): got %t, %v; want true, nil"
		t.Errorf(tmpl, username, avail, err)
	}
}
