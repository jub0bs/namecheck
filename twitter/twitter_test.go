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
