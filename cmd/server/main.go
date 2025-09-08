package main

import (
	"fmt"
	"log"
	"net/http"
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
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /check", handleCheck)
	if err := http.ListenAndServe(":8080", mux); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func handleCheck(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	const n = 20
	checkers := make([]Checker, n)
	gh := github.GitHub{Client: http.DefaultClient}
	for i := range n {
		checkers[i] = &gh
	}
	resultCh := make(chan Result)
	errorCh := make(chan error)
	var wg sync.WaitGroup
	for _, checker := range checkers {
		wg.Add(1)
		go check(checker, username, &wg, resultCh, errorCh)
	}
	go func() {
		wg.Wait()
		close(resultCh)
	}()
	var results []Result
	var finished bool
	for !finished {
		select {
		case <-errorCh:
			w.WriteHeader(http.StatusInternalServerError)
			return
		case res, ok := <-resultCh:
			if !ok {
				finished = true
				continue
			}
			results = append(results, res)
		}
	}
	fmt.Fprint(w, results)
}

func check(
	checker Checker,
	username string,
	wg *sync.WaitGroup,
	resultCh chan<- Result,
	errorCh chan<- error,
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
	avail, err := checker.IsAvailable(username)
	if err != nil {
		errorCh <- err
		return
	}
	res.Available = avail
	resultCh <- res
}
