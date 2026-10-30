package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/jub0bs/namecheck/github"
)

type Result struct {
	Platform  string
	Valid     bool
	Available bool
	Err       error
}

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
	resultCh := make(chan Result)
	var wg sync.WaitGroup
	for _, checker := range checkers {
		wg.Go(func() { check(checker, username, resultCh) })
	}
	go func() {
		wg.Wait()
		close(resultCh)
	}()
	var results []Result
	for res := range resultCh {
		results = append(results, res)
	}
	fmt.Println(results)
}

func check(checker Checker, username string, resultCh chan<- Result) {
	res := Result{
		Platform: checker.String(),
		Valid:    checker.IsValid(username),
	}
	if !res.Valid {
		resultCh <- res
		return
	}
	res.Available, res.Err = checker.IsAvailable(username)
	resultCh <- res
}
