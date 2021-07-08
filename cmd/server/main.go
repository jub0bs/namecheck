package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/jub0bs/namecheck"
	"github.com/jub0bs/namecheck/github"
	"github.com/jub0bs/namecheck/twitter"
)

type Result struct {
	Username  string
	Platform  string
	Valid     bool
	Available bool
}

func main() {
	// Hello world, the web server

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/check", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello, world!\n")
}

func handler(w http.ResponseWriter, req *http.Request) {
	username := req.URL.Query().Get("username")
	tw := twitter.Twitter{
		Client: http.DefaultClient,
	}
	gh := github.GitHub{
		Client: http.DefaultClient,
	}
	// let's pretend that we support many platforms
	var checkers []namecheck.Checker
	for i := 0; i < 2; i++ {
		checkers = append(checkers, &tw, &gh)
	}

	ch := make(chan Result, len(checkers))
	errc := make(chan error, len(checkers))
	var wg sync.WaitGroup
	wg.Add(len(checkers))
	for _, checker := range checkers {
		go check(checker, username, &wg, ch, errc)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	results := make([]Result, 0, len(checkers))
	var done bool
	for !done {
		select {
		case res, ok := <-ch:
			if !ok {
				done = true
				continue
			}
			results = append(results, res)
		case err := <-errc:
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}

	fmt.Fprint(w, results)
}

func check(
	checker namecheck.Checker,
	username string,
	wg *sync.WaitGroup,
	ch chan<- Result,
	errc chan<- error,
) {
	defer wg.Done()
	res := Result{
		Username: username,
		Platform: checker.String(),
	}
	res.Valid = checker.IsValid(username)
	if !res.Valid {
		ch <- res
		return
	}
	avail, err := checker.IsAvailable(username)
	if err != nil {
		errc <- err
		return
	}
	res.Available = avail
	ch <- res
}
