package twitter_test

import (
	"testing"

	"github.com/jub0bs/namecheck/twitter"
)

func TestUsernameTooLong(t *testing.T) {
	var tw twitter.Twitter
	username := "obviously_longer_than_15_chars"
	want := false
	got := tw.IsValid(username)
	if got != want {
		t.Errorf("twitter.IsValid(%q): got %t; want %t", username, got, want)
	}
}

func TestUsernameTooShort(t *testing.T) {
	var tw twitter.Twitter
	username := "foo"
	want := false
	got := tw.IsValid(username)
	if got != want {
		t.Errorf("twitter.IsValid(%q): got %t; want %t", username, got, want)
	}
}

func TestUsernameContainsTwitter(t *testing.T) {
	var tw twitter.Twitter
	username := "jub0bsOnTwitter"
	want := false
	got := tw.IsValid(username)
	if got != want {
		t.Errorf("twitter.IsValid(%q): got %t; want %t", username, got, want)
	}
}

func TestUsernameContainsIllegalChars(t *testing.T) {
	var tw twitter.Twitter
	username := "-jub-0bs-"
	want := false
	got := tw.IsValid(username)
	if got != want {
		t.Errorf("twitter.IsValid(%q): got %t; want %t", username, got, want)
	}
}

func TestUsernameIsValid(t *testing.T) {
	var tw twitter.Twitter
	username := "jub0bs"
	want := true
	got := tw.IsValid(username)
	if got != want {
		t.Errorf("twitter.IsValid(%q): got %t; want %t", username, got, want)
	}
}
