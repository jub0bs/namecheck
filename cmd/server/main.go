package main

import (
	"encoding/json/v2"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/jub0bs/namecheck/github"
)

type Result struct {
	Platform  string `json:"platform"`
	Valid     bool   `json:"valid"`
	Available bool   `json:"available"`
	Err       error  `json:"error,omitempty"`
}

type Checker interface {
	IsValid(string) bool
	IsAvailable(string) (bool, error)
	fmt.Stringer
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
	w.Header().Set("Content-Type", "application/json")
	data := struct {
		Username string   `json:"username"`
		Results  []Result `json:"results,omitempty"`
	}{
		Username: username,
		Results:  results,
	}
	if err := json.MarshalWrite(w, data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
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
