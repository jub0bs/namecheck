package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/jub0bs/namecheck/github"
)

type Checker interface {
	IsValid(string) bool
	IsAvailable(string) (bool, error)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <username>\n", os.Args[0])
		os.Exit(2)
	}
	username := os.Args[1]
	gh := github.GitHub{
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
	const n = 20
	checkers := make([]Checker, 0, n)
	for range n {
		checkers = append(checkers, &gh)
	}
	var wg sync.WaitGroup
	wg.Add(len(checkers))
	for _, checker := range checkers {
		go check(checker, username, &wg)
	}
	wg.Wait()
}

func check(checker Checker, username string, wg *sync.WaitGroup) {
	defer wg.Done()
	valid := checker.IsValid(username)
	fmt.Printf("validity of %q on %s: %t\n", username, checker, valid)
	if !valid {
		return
	}
	avail, err := checker.IsAvailable(username)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("availability of %q on %s: %t\n", username, checker, avail)
	}
}
