package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/jub0bs/namecheck/bluesky"
	"github.com/jub0bs/namecheck/github"
)

type Checker interface {
	IsValid(string) bool
	IsAvailable(string) (bool, error)
	fmt.Stringer
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: namecheck <username>")
		os.Exit(1)
	}
	username := os.Args[1]
	gh := github.GitHub{
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
	var bs bluesky.Bluesky
	checkers := []Checker{&gh, &bs}
	for _, checker := range checkers {
		valid := checker.IsValid(username)
		fmt.Printf("validity of %q on %s: %t\n", username, checker, valid)
		if !valid {
			continue
		}
		available, err := checker.IsAvailable(username)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("availability of %q on %s: %t\n", username, checker, available)
	}
}
