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
	"golang.org/x/sync/semaphore"
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
	mu.Lock()
	stats := maps.Clone(stats) // deep copy
	mu.Unlock()
	if err := json.MarshalWrite(w, stats); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type token struct{}

var sem = semaphore.NewWeighted(16)

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
		Client: http.DefaultClient,
	}
	var checkers []Checker
	for range 20 {
		checkers = append(checkers, &gh)
	}
	resultCh := make(chan Result)
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	go func() {
		for _, checker := range checkers {
			sem.Acquire(context.Background(), 1)
			wg.Go(func() {
				defer func() { sem.Release(1) }()
				check(ctx, username, checker, resultCh)
			})
		}
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

func check(
	ctx context.Context,
	username string,
	checker Checker,
	resultCh chan<- Result,
) {
	res := Result{
		Platform: checker.String(),
		Valid:    checker.IsValid(username),
	}
	if !res.Valid {
		select {
		case <-ctx.Done():
			return
		case resultCh <- res:
		}
		return
	}
	res.Available, res.Err = checker.IsAvailable(ctx, username)
	select {
	case <-ctx.Done():
		return
	case resultCh <- res:
	}
}
