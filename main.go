package main

import (
	"fmt"
	"log"
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
	checkers := []Checker{
		&github.GitHub{Client: http.DefaultClient},
		&bluesky.Bluesky{},
	}
	for _, checker := range checkers {
		valid := checker.IsValid(username)
		fmt.Printf("validity of %q on GitHub: %t\n", username, valid)
		if valid {
			avail, err := checker.IsAvailable(username)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("availability of %q on GitHub: %t\n", username, avail)
		}
	}
}
