package main

import (
	"fmt"
	"log"
	"net/http"
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
	http.HandleFunc("GET /check", handleCheck)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleCheck(w http.ResponseWriter, req *http.Request) {
	username := req.URL.Query().Get("username")
	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
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
	fmt.Fprint(w, results)
}

func check(
	checker Checker,
	username string,
	wg *sync.WaitGroup,
	resultCh chan<- Result,
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
