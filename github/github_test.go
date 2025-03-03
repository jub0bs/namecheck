package github_test

import (
	"strings"
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

func TestUsernameStartsWithAHyphen(t *testing.T) {
	const username = "-jub0bs-on-GitHub"
	const want = false
	got := github.IsValid(username)
	if got != want {
		t.Errorf("github.IsValid(%q): got %t; want %t", username, got, want)
	}
}

func TestUsernameEndsWithAHyphen(t *testing.T) {
	const username = "jub0bs-on-GitHub-"
	const want = false
	got := github.IsValid(username)
	if got != want {
		t.Errorf("github.IsValid(%q): got %t; want %t", username, got, want)
	}
}

func TestUsernameTooShort(t *testing.T) {
	const username = "ab"
	const want = false
	got := github.IsValid(username)
	if got != want {
		t.Errorf("github.IsValid(%q): got %t; want %t", username, got, want)
	}
}

func TestUsernameTooLong(t *testing.T) {
	username := strings.Repeat("a", 40)
	const want = false
	got := github.IsValid(username)
	if got != want {
		t.Errorf("github.IsValid(%q): got %t; want %t", username, got, want)
	}
}

func TestUsernameContainsInvalidChars(t *testing.T) {
	const username = "ju^b0bs"
	const want = false
	got := github.IsValid(username)
	if got != want {
		t.Errorf("github.IsValid(%q): got %t; want %t", username, got, want)
	}
}

func TestUsernameAllGood(t *testing.T) {
	const username = "jub0bs"
	const want = true
	got := github.IsValid(username)
	if got != want {
		t.Errorf("github.IsValid(%q): got %t; want %t", username, got, want)
	}
}
