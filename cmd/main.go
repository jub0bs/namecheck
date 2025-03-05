package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"

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
	var checkers []Checker
	for range 20 {
		checkers = append(checkers, &gh)
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
	valid := checker.IsValid(username)
	fmt.Printf("validity of %q on %s: %t\n", username, checker, valid)
	if !valid {
		return
	}
	avail, err := checker.IsAvailable(username)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Printf("available of %q on %s: %t\n", username, checker, avail)
}
