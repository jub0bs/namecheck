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
		"too short":                        {"ab", false},
		"too long":                         {strings.Repeat("a", 40), false},
		"contains invalid chars":           {"^_^'", false},
		"all good":                         {"jub0bs", true},
		// other test cases...
	}
	gh := github.GitHub{
		Client: http.DefaultClient,
	}
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

func TestIsAvailable200OK(t *testing.T) {
	gh := github.GitHub{
		Client: &stub.SuccessfulGetter{StatusCode: http.StatusOK},
	}
	avail, err := gh.IsAvailable("whatever")
	if err != nil || avail {
		t.Errorf("IsAvailable(...): got %t, %v; want false, nil", avail, err)
	}
}

func TestIsAvailable404NotFound(t *testing.T) {
	gh := github.GitHub{
		Client: &stub.SuccessfulGetter{StatusCode: http.StatusNotFound},
	}
	avail, err := gh.IsAvailable("whatever")
	if err != nil || !avail {
		t.Errorf("IsAvailable(...): got %t, %v; want true, nil", avail, err)
	}
}
