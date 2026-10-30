package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

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
	var checkers []Checker
	const n = 20
	for range n {
		checkers = append(checkers, &gh)
	}
	var wg sync.WaitGroup
	for _, checker := range checkers {
		wg.Go(func() { check(checker, username) })
	}
	wg.Wait()
}

func check(checker Checker, username string) {
	valid := checker.IsValid(username)
	fmt.Printf("validity of %q on %s: %t\n", username, checker, valid)
	if !valid {
		return
	}
	available, err := checker.IsAvailable(username)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("availability of %q on %s: %t\n", username, checker, available)
}
