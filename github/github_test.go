package github_test // external test package: black-box testing

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
		{"contains two consecutive hyphens", "jub0bs--on-GitHub", false},
		{"too short", "ab", false},
		{"too long", strings.Repeat("a", 40), false},
		{"starts with a hyphen", "-jub0bs", false},
		{"ends with a hyphen", "jub0bs-", false},
		{"contains illegal chars", "jub&obs", false},
		{"all good", "jub0bs", true},
		// other test cases...
	}
	for _, tc := range testCases {
		f := func(t *testing.T) {
			got := github.IsValid(tc.username)
			if got != tc.want {
				const tmpl = "github.IsValid(%q): got %t; want %t"
				t.Errorf(tmpl, tc.username, got, tc.want)
			}
		}
		t.Run(tc.desc, f)
	}
}
