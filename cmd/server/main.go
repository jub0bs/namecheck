package main

import (
	"context"
	"encoding/json/v2"
	"fmt"
	"log"
	"maps"
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
	IsAvailable(context.Context, string) (bool, error)
	fmt.Stringer
}

var (
	stats = make(map[string]int) // guarded by mu
	mu    sync.Mutex
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
	if err := http.ListenAndServe(":8080", corsMw.Wrap(mux)); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func handleStats(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	stats := maps.Clone(stats)
	mu.Unlock()
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
	mu.Lock()
	stats[username]++
	mu.Unlock()
	gh := github.GitHub{
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
	var checkers []Checker
	const n = 80
	for range n {
		checkers = append(checkers, &gh)
	}
	resultCh := make(chan Result)
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()
	for _, checker := range checkers {
		wg.Go(func() { check(ctx, checker, username, resultCh) })
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

func check(ctx context.Context, checker Checker, username string, resultCh chan<- Result) {
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
	res.Available, res.Err = checker.IsAvailable(ctx, username)
	select {
	case <-ctx.Done():
		break
	case resultCh <- res:
	}
}
