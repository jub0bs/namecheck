package github_test

import (
	"strings"
	"testing"

	"github.com/jub0bs/namecheck/github"
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
				t.Errorf("%s: github.IsValid(%q): got %t; want %t", tc.desc, tc.username, got, tc.want)
			}
		}
		t.Run(tc.desc, f)
	}
}
