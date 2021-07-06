package github_test

import (
	"testing"

	"github.com/jub0bs/namecheck/github"
)

func TestUsernameTooLong(t *testing.T) {
	username := "obviously_longer_than_39_chars_haaaaaaaaaaaaaaaaaaaaaaaaa"
	want := false
	got := github.IsValid(username)
	if got != want {
		t.Errorf("github.IsValid(%q): got %t; want %t", username, got, want)
	}
}

func TestUsernameTooShort(t *testing.T) {
	username := "vi"
	want := false
	got := github.IsValid(username)
	if got != want {
		t.Errorf("github.IsValid(%q): got %t; want %t", username, got, want)
	}
}

func TestUsernameContainsIllegalChars(t *testing.T) {
	username := "_jub0bs_"
	want := false
	got := github.IsValid(username)
	if got != want {
		t.Errorf("github.IsValid(%q): got %t; want %t", username, got, want)
	}
}

func TestUsernameContainsIllegalPattern(t *testing.T) {
	username := "-jub0--bs-"
	want := false
	got := github.IsValid(username)
	if got != want {
		t.Errorf("github.IsValid(%q): got %t; want %t", username, got, want)
	}
}

func TestUsernameIsValid(t *testing.T) {
	username := "jub0bs"
	want := true
	got := github.IsValid(username)
	if got != want {
		t.Errorf("github.IsValid(%q): got %t; want %t", username, got, want)
	}
}
