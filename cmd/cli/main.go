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
		fmt.Fprint(os.Stderr, "usage: namecheck <username>\n")
		os.Exit(1)
	}
	username := os.Args[1]
	gh := github.GitHub{Client: http.DefaultClient}
	const n = 16
	checkers := make([]Checker, n)
	for i := range n {
		checkers[i] = &gh
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
	resultCh chan Result,
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
