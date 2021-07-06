package twitter_test

import (
	"testing"

	"github.com/jub0bs/namecheck/twitter"
)

func TestUsernameTooLong(t *testing.T) {
	username := "obviously_longer_than_15_chars"
	want := false
	got := twitter.IsValid(username)
	if got != want {
		t.Errorf("twitter.IsValid(%q): got %t; want %t", username, got, want)
	}
}

func TestUsernameTooShort(t *testing.T) {
	username := "foo"
	want := false
	got := twitter.IsValid(username)
	if got != want {
		t.Errorf("twitter.IsValid(%q): got %t; want %t", username, got, want)
	}
}

func TestUsernameContainsTwitter(t *testing.T) {
	username := "jub0bsOnTwitter"
	want := false
	got := twitter.IsValid(username)
	if got != want {
		t.Errorf("twitter.IsValid(%q): got %t; want %t", username, got, want)
	}
}

func TestUsernameContainsIllegalChars(t *testing.T) {
	username := "-jub-0bs-"
	want := false
	got := twitter.IsValid(username)
	if got != want {
		t.Errorf("twitter.IsValid(%q): got %t; want %t", username, got, want)
	}
}

func TestUsernameIsValid(t *testing.T) {
	username := "jub0bs"
	want := true
	got := twitter.IsValid(username)
	if got != want {
		t.Errorf("twitter.IsValid(%q): got %t; want %t", username, got, want)
	}
}
