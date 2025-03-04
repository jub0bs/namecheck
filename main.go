package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jub0bs/namecheck/bluesky"
	"github.com/jub0bs/namecheck/github"
)

type Checker interface {
	IsValid(string) bool
	IsAvailable(string) (bool, error)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <username>\n", os.Args[0])
		os.Exit(1)
	}
	username := os.Args[1]
	gh := github.GitHub{
		Client: http.DefaultClient,
	}
	var bs bluesky.Bluesky
	checkers := []Checker{&gh, &bs}
	for _, checker := range checkers {
		valid := checker.IsValid(username)
		fmt.Printf("validity of %q %T: %t\n", username, checker, valid)
		if !valid {
			continue
		}
		avail, err := checker.IsAvailable(username)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		fmt.Printf("available of %q %T: %t\n", username, checker, avail)
	}
}
