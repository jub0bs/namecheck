package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jub0bs/namecheck/github"
)

type Checker interface {
	IsValid(string) bool
	IsAvailable(string) (bool, error)
	fmt.Stringer
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, "usage: namecheck <username>\n")
		os.Exit(1)
	}
	username := os.Args[1]
	gh := github.GitHub{Client: http.DefaultClient}
	const n = 64
	checkers := make([]Checker, n)
	for i := range n {
		checkers[i] = &gh
	}
	for _, checker := range checkers {
		if !checker.IsValid(username) {
			continue
		}
		avail, err := checker.IsAvailable(username)
		if err != nil {
			log.Fatal(err)
		}
		if !avail {
			continue
		}
		fmt.Printf("%q is valid and available on %s\n", username, checker)
	}
}
