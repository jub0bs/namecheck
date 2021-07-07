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
	username := os.Args[1]
	var (
		tw twitter.Twitter
		gh github.GitHub
	)
	fmt.Println(&tw, &gh)
	var (
		valid bool
		avail bool
		err   error
	)

	valid = tw.IsValid(username)
	if valid {
		avail, err = tw.IsAvailable(username)
	}
	fmt.Printf(
		"[%q on %s] valid: %t; available: %t; error: %v\n",
		username,
		"Twitter", valid,
		avail,
		err,
	)
	valid = gh.IsValid(username)
	if valid {
		avail, err = gh.IsAvailable(username)
	}
	fmt.Printf(
		"[%q on %s] valid: %t; available: %t; error: %v\n",
		username,
		"GitHub", valid,
		avail,
		err,
	)
}
