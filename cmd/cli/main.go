package main

import (
	"fmt"
	"os"

	"github.com/jub0bs/namecheck"
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
	// let's pretend that we support many platforms
	var checkers []namecheck.Checker
	for i := 0; i < 50; i++ {
		checkers = append(checkers, &tw, &gh)
	}
	for _, checker := range checkers {
		valid := checker.IsValid(username)
		var (
			avail bool
			err   error
		)
		if valid {
			avail, err = checker.IsAvailable(username)
		}
		fmt.Printf(
			"[%q on %s] valid: %t; available: %t; error: %v\n",
			username,
			checker.String(),
			valid,
			avail,
			err,
		)
	}
}
