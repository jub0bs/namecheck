package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jub0bs/namecheck/github"
	"github.com/jub0bs/namecheck/reddit"
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
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: namecheck <username>")
		os.Exit(1)
	}
	username := os.Args[1]
	gh := github.GitHub{
		Client: http.DefaultClient,
	}
	var re reddit.Reddit
	checkers := []Checker{&gh, &re}
	for _, checker := range checkers {
		valid := checker.IsValid(username)
		fmt.Printf("%q is valid on %s: %t\n", username, checker, valid)
		if valid {
			avail, err := checker.IsAvailable(username)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("%q is available on %s: %t\n", username, checker, avail)
			}
		}
	}
}
