package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/jub0bs/cors"
	"github.com/jub0bs/namecheck/bluesky"
	"github.com/jub0bs/namecheck/github"
)

type Checker interface {
	IsValid(string) bool
	IsAvailable(string) (bool, error)
	fmt.Stringer
}

type Result struct {
	Platform  string `json:"platform"`
	Valid     bool   `json:"valid"`
	Available bool   `json:"available"`
	Err       error  `json:"error,omitempty"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /check", handleCheck)
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

func handleCheck(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	gh := github.GitHub{Client: http.DefaultClient}
	bs := bluesky.Bluesky{}
	var checkers []Checker
	for range 1 {
		checkers = append(checkers, &gh, &bs)
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
	data := struct {
		Username string   `json:"username"`
		Results  []Result `json:"results,omitempty"`
	}{
		Username: username,
		Results:  results,
	}
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if err := enc.Encode(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
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
