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
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <username>\n", os.Args[0])
		os.Exit(2)
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
		fmt.Printf("validity of %q on ?: %t\n", username, valid)
		if valid {
			avail, err := checker.IsAvailable(username)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("availability of %q on ?: %t\n", username, avail)
			}
		}
	}
}
