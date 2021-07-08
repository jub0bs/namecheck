package github_test

import (
	"net/http"
	"testing"

	"github.com/jub0bs/namecheck/github"
	"github.com/jub0bs/namecheck/stub"
)

func TestUsernameTooLong(t *testing.T) {
	var gh github.GitHub
	username := "obviously_longer_than_39_chars_haaaaaaaaaaaaaaaaaaaaaaaaa"
	want := false
	got := gh.IsValid(username)
	if got != want {
		t.Errorf("github.IsValid(%q): got %t; want %t", username, got, want)
	}
}

func TestUsernameTooShort(t *testing.T) {
	var gh github.GitHub
	username := "vi"
	want := false
	got := gh.IsValid(username)
	if got != want {
		t.Errorf("github.IsValid(%q): got %t; want %t", username, got, want)
	}
}

func TestUsernameContainsIllegalChars(t *testing.T) {
	var gh github.GitHub
	username := "_jub0bs_"
	want := false
	got := gh.IsValid(username)
	if got != want {
		t.Errorf("github.IsValid(%q): got %t; want %t", username, got, want)
	}
}

func TestUsernameContainsIllegalPattern(t *testing.T) {
	var gh github.GitHub
	username := "-jub0--bs-"
	want := false
	got := gh.IsValid(username)
	if got != want {
		t.Errorf("github.IsValid(%q): got %t; want %t", username, got, want)
	}
}

func TestUsernameIsValid(t *testing.T) {
	var gh github.GitHub
	username := "jub0bs"
	want := true
	got := gh.IsValid(username)
	if got != want {
		t.Errorf("github.IsValid(%q): got %t; want %t", username, got, want)
	}
}

func TestUsernameIsAvailable404(t *testing.T) {
	gh := github.GitHub{
		Client: &stub.SuccessfulClient{StatusCode: http.StatusNotFound},
	}
	username := "whatever"
	avail, err := gh.IsAvailable(username)
	if err != nil || !avail {
		t.Errorf(
			"github.IsAvailable(%q) = %t, %v; want true, nil",
			username,
			avail,
			err,
		)
	}
}
