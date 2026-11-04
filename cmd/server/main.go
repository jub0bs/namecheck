package main

import (
	"encoding/json/v2"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/jub0bs/cors"
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

var stats = make(map[string]uint)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /check", handleCheck)
	mux.HandleFunc("GET /stats", handleStats)
	corsMw, err := cors.NewMiddleware(cors.Config{
		Origins: []string{"https://namecheck.jub0bs.dev"},
	})
	if err != nil {
		log.Fatal(err)
	}
	handler := corsMw.Wrap(mux)
	if err := http.ListenAndServe(":8080", handler); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func handleStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.MarshalWrite(w, stats); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func handleCheck(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	stats[username]++
	gh := github.GitHub{
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
	var checkers []Checker
	for range 20 {
		checkers = append(checkers, &gh)
	}
	resultCh := make(chan Result)
	var wg sync.WaitGroup
	for _, checker := range checkers {
		wg.Go(func() { check(username, checker, resultCh) })
	}
	go func() {
		wg.Wait()
		close(resultCh)
	}()
	var results []Result
	for res := range resultCh {
		results = append(results, res)
	}
	data := struct {
		Username string   `json:"username"`
		Results  []Result `json:"results,omitempty"`
	}{
		Username: username,
		Results:  results,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.MarshalWrite(w, data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func check(username string, checker Checker, resultCh chan<- Result) {
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
