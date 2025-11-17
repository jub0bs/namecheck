package github_test

import (
	"testing"

	"github.com/jub0bs/namecheck/github"
)

func TestUsernameContainsTwoConsecutiveHyphens(t *testing.T) {
	const username = "jub0bs--on-GitHub"
	const want = false
	got := github.IsValid(username)
	if got != want {
		t.Errorf("github.IsValid(%q): got %t; want %t", username, got, want)
	}
}
