package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jub0bs/namecheck/bluesky"
	"github.com/jub0bs/namecheck/github"
)

type Validator interface {
	IsValid(string) bool
}

type Availabler interface {
	IsAvailable(string) (bool, error)
}

type Checker interface {
	Validator
	Availabler
	fmt.Stringer
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
		fmt.Printf("validity on %s: %t\n", checker, valid)
		if !valid {
			continue
		}
		avail, err := checker.IsAvailable(username)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Printf("availability on %s: %t\n", checker, avail)
	}
}
