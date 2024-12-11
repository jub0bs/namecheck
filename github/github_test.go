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
		desc     string
		username string
		want     bool
	}
	testCases := []TestCase{
		{desc: "contains two consecutive hyphens", username: "jub0bs--on-GitHub"},
		{"starts with a hyphen", "-jub0bs-on-GitHub", false},
		{"ends with a hyphen", "jub0bs-on-GitHub-", false},
		{"too long", strings.Repeat("a", 40), false},
		{"too short", "ab", false},
		{"contains illegal chars", "a^b", false},
		{"all good", "jub0bs", true},
		// other test cases...
	}
	var gh github.GitHub
	for _, tc := range testCases {
		f := func(t *testing.T) {
			got := gh.IsValid(tc.username)
			if got != tc.want {
				t.Errorf("github.IsValid(%q): got %t; want %t", tc.username, got, tc.want)
			}
		}
		t.Run(tc.desc, f)
	}
}

func TestIsAvailable200OK(t *testing.T) {
	gh := github.GitHub{
		Client: &stub.SuccessfulClient{StatusCode: http.StatusOK},
	}
	avail, err := gh.IsAvailable("whatever")
	if err != nil || avail {
		const tmpl = "github.IsAvailable(): got %t, %v; want false, nil"
		t.Errorf(tmpl, avail, err)
	}
}

func TestIsAvailable404NotFound(t *testing.T) {
	gh := github.GitHub{
		Client: &stub.SuccessfulClient{StatusCode: http.StatusNotFound},
	}
	avail, err := gh.IsAvailable("whatever")
	if err != nil || !avail {
		const tmpl = "github.IsAvailable(): got %t, %v; want true, nil"
		t.Errorf(tmpl, avail, err)
	}
}

func TestIsAvailableOtherStatusCode(t *testing.T) {
	gh := github.GitHub{
		Client: &stub.SuccessfulClient{StatusCode: http.StatusServiceUnavailable},
	}
	avail, err := gh.IsAvailable("whatever")
	if err == nil || avail {
		const tmpl = "github.IsAvailable(): got %t, %v; want false, some non-nil error"
		t.Errorf(tmpl, avail, err)
	}
}
