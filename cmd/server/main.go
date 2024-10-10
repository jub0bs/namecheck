package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/jub0bs/cors"
	"github.com/jub0bs/namecheck"
	"github.com/jub0bs/namecheck/github"
	"github.com/jub0bs/namecheck/reddit"
)

type Result struct {
	Username  string `json:"username"`
	Platform  string `json:"platform"`
	Valid     bool   `json:"valid"`
	Available bool   `json:"available"`
}

var m = make(map[string]uint)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /check", handleCheck)
	mux.HandleFunc("GET /stats", handleStats)

	// instantiate a CORS middleware whose config suits your needs
	corsMw, err := cors.NewMiddleware(cors.Config{
		Origins: []string{"*"},
	})
	if err != nil {
		log.Fatal(err)
	}

	// apply the CORS middleware
	handler := corsMw.Wrap(mux)

	// start the server on port 8080; make sure to pass your custom handler
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}

func handleStats(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if err := enc.Encode(m); err != nil {
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
	m[username]++
	gh := github.GitHub{
		Client: http.DefaultClient,
	}
	var re reddit.Reddit
	const n = 40
	checkers := make([]namecheck.Checker, 0, 2*n)
	for range n {
		checkers = append(checkers, &gh, &re)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	resultCh := make(chan Result)
	errCh := make(chan error)
	var wg sync.WaitGroup
	for _, checker := range checkers {
		wg.Add(1)
		go check(ctx, checker, username, &wg, resultCh, errCh)
	}
	go func() {
		wg.Wait()
		close(resultCh)
	}()
	var results []Result
	var finished bool
	for !finished {
		select {
		case <-errCh:
			cancel()
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
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if err := enc.Encode(results); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func check(
	ctx context.Context,
	checker namecheck.Checker,
	username string,
	wg *sync.WaitGroup,
	resultCh chan<- Result,
	errCh chan<- error,
) {
	defer wg.Done()
	res := Result{
		Username: username,
		Platform: checker.String(),
		Valid:    checker.IsValid(username),
	}
	if !res.Valid {
		select {
		case <-ctx.Done():
		case resultCh <- res:
		}
		return
	}
	avail, err := checker.IsAvailable(username)
	if err != nil {
		select {
		case <-ctx.Done():
		case errCh <- err:
		}
		return
	}
	res.Available = avail
	select {
	case <-ctx.Done():
	case resultCh <- res:
	}
}
