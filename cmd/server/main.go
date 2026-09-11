package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"maps"
	"net/http"
	"sync"

	"github.com/jub0bs/cors"
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

var (
	stats = make(map[string]uint)
	mu    sync.Mutex // guards stats
)

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
	enc := json.NewEncoder(w)
	mu.Lock()
	stats := maps.Clone(stats)
	mu.Unlock()
	if err := enc.Encode(stats); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func handleCheck(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	mu.Lock()
	stats[username]++
	mu.Unlock()
	gh := github.GitHub{Client: http.DefaultClient}
	const n = 16
	checkers := make([]Checker, n)
	for i := range n {
		checkers[i] = &gh
	}
	resultCh := make(chan Result)
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for _, checker := range checkers {
		wg.Add(1)
		go check(ctx, checker, username, &wg, resultCh)
	}
	go func() {
		wg.Wait()
		close(resultCh)
	}()
	var results []Result
	for res := range resultCh {
		if res.Err != nil {
			cancel()
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		results = append(results, res)
	}
	type respBody struct {
		Username string   `json:"username"`
		Results  []Result `json:"results,omitempty"`
	}
	rb := respBody{
		Username: username,
		Results:  results,
	}
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if err := enc.Encode(rb); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func check(
	ctx context.Context,
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
		select {
		case <-ctx.Done():
			break
		case resultCh <- res:
		}
		return
	}
	res.Available, res.Err = checker.IsAvailable(username)
	select {
	case <-ctx.Done():
		break
	case resultCh <- res:
	}
}
