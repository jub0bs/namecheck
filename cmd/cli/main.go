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
	var valid []string
	var invalid []string
	for _, username := range os.Args[1:] {
		if twitter.IsValid(username) && github.IsValid(username) {
			valid = append(valid, username)
		} else {
			invalid = append(invalid, username)
		}
	}
	fmt.Println("usernames valid on both Twitter and GitHub:", valid)
	fmt.Println("usernames invalid on either Twitter or GitHub:", invalid)
}
