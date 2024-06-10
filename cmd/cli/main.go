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

type Result struct {
	Username  string
	Platform  string
	Valid     bool
	Available bool
	Err       error
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
	resultCh := make(chan Result)
	var wg sync.WaitGroup
	for _, checker := range checkers {
		wg.Add(1)
		go check(checker, username, &wg, resultCh)
	}
	go func() {
		wg.Wait()
		close(resultCh)
	}()
	results := make([]Result, 0, len(checkers))
	for res := range resultCh {
		results = append(results, res)
	}
	fmt.Println(results)
}

func check(
	checker Checker,
	username string,
	wg *sync.WaitGroup,
	resultCh chan Result,
) {
	defer wg.Done()
	res := Result{
		Username: username,
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
