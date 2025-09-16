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
	"github.com/jub0bs/namecheck/bluesky"
	"github.com/jub0bs/namecheck/github"
)

type Checker interface {
	IsValid(string) bool
	IsAvailable(context.Context, string) (bool, error)
	fmt.Stringer
}

type Result struct {
	Platform  string `json:"platform"`
	Valid     bool   `json:"valid"`
	Available bool   `json:"available"`
}

var (
	stats = make(map[string]uint) // guarded by mu
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
	handler := corsMw.Wrap(mux)
	if err := http.ListenAndServe(":8080", handler); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func handleStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	mu.Lock()
	statsCopy := maps.Clone(stats)
	mu.Unlock()
	if err := enc.Encode(statsCopy); err != nil {
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
	gh := github.GitHub{Client: http.DefaultClient}
	bs := bluesky.Bluesky{}
	var checkers []Checker
	for range 40 {
		checkers = append(checkers, &gh, &bs)
	}
	resultCh := make(chan Result)
	errorCh := make(chan error)
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()
	var wg sync.WaitGroup
	for _, checker := range checkers {
		wg.Add(1)
		go check(ctx, checker, username, &wg, resultCh, errorCh)
	}
	go func() {
		wg.Wait()
		close(resultCh)
	}()
	var results []Result
	var finished bool
	for !finished {
		select {
		case err := <-errorCh:
			cancel()
			fmt.Println(err)
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
	ctx context.Context,
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
		send(ctx, resultCh, res)
		return
	}
	avail, err := checker.IsAvailable(ctx, username)
	if err != nil {
		send(ctx, errorCh, err)
		return
	}
	res.Available = avail
	send(ctx, resultCh, res)
}

func send[T any](ctx context.Context, ch chan<- T, v T) {
	select {
	case <-ctx.Done():
	case ch <- v:
	}
}
