package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"

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
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: namecheck <username>")
		os.Exit(1)
	}
	username := os.Args[1]
	gh := github.GitHub{
		Client: http.DefaultClient,
	}
	re := reddit.Reddit{
		Client: http.DefaultClient,
	}
	var checkers []Checker
	for range 20 {
		checkers = append(checkers, &gh, &re)
	}
	var wg sync.WaitGroup
	for _, checker := range checkers {
		wg.Add(1)
		go check(checker, username, &wg)
	}
	wg.Wait()
}

func check(checker Checker, username string, wg *sync.WaitGroup) {
	defer wg.Done()
	if !checker.IsValid(username) {
		return
	}
	avail, err := checker.IsAvailable(username)
	if err != nil || !avail {
		return
	}
	fmt.Printf("%q is valid and available on %s\n", username, checker)
}
