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

type Result struct {
	Platform  string
	Valid     bool
	Available bool
	Err       error
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
	const n = 20
	checkers := make([]Checker, 0, n)
	for range n {
		checkers = append(checkers, &gh)
	}
	resultCh := make(chan Result)
	var wg sync.WaitGroup
	wg.Add(len(checkers))
	for _, checker := range checkers {
		go check(checker, username, &wg, resultCh)
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

func check(
	checker Checker,
	username string,
	wg *sync.WaitGroup,
	resultCh chan<- Result,
) {
	defer wg.Done()
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
