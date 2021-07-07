package main

import (
	"fmt"
	"os"

	"github.com/jub0bs/namecheck/github"
	"github.com/jub0bs/namecheck/twitter"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: namecheck <username>...")
		os.Exit(1)
	}
	var (
		tw      twitter.Twitter
		gh      github.GitHub
		valid   []string
		invalid []string
	)
	for _, username := range os.Args[1:] {
		if tw.IsValid(username) && gh.IsValid(username) {
			valid = append(valid, username)
		} else {
			invalid = append(invalid, username)
		}
	}
	fmt.Println("usernames valid on both Twitter and GitHub:", valid)
	fmt.Println("usernames invalid on either Twitter or GitHub:", invalid)
}
