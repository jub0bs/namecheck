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
	Username  string `json:"user_name"`
	Platform  string `json:"platform"`
	Valid     bool   `json:"valid"`
	Available bool   `json:"available"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /check", handleCheck)
	corsMw, err := cors.NewMiddleware(cors.Config{
		Origins: []string{"*"},
	})
	if err != nil {
		log.Fatal(err)
	}
	handler := corsMw.Wrap(mux)
	log.Fatal(http.ListenAndServe(":8080", handler))
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
	var checkers []namecheck.Checker
	for range 20 {
		checkers = append(checkers, &gh, &re)
	}
	resultCh := make(chan Result)
	errorCh := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())
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
	results := make([]Result, 0, len(checkers))
	var finished bool
	for !finished {
		select {
		case <-errorCh:
			w.WriteHeader(http.StatusInternalServerError)
			cancel()
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
	}
}

func check(
	ctx context.Context,
	checker namecheck.Checker,
	username string,
	wg *sync.WaitGroup,
	resultCh chan<- Result,
	errorCh chan<- error,
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
		case errorCh <- err:
		}
		return
	}
	res.Available = avail
	select {
	case <-ctx.Done():
	case resultCh <- res:
	}
}
