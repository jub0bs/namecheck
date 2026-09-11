package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/jub0bs/namecheck/github"
)

type Checker interface {
	IsValid(string) bool
	IsAvailable(string) (bool, error)
	fmt.Stringer
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, "usage: namecheck <username>\n")
		os.Exit(1)
	}
	username := os.Args[1]
	gh := github.GitHub{Client: http.DefaultClient}
	const n = 64
	checkers := make([]Checker, n)
	for i := range n {
		checkers[i] = &gh
	}
	var wg sync.WaitGroup
	for _, checker := range checkers {
		wg.Add(1)
		go check(checker, username, &wg)
	}
	wg.Wait()
}

func check(checker Checker, username string, wg *sync.WaitGroup) {
	defer wg.Done()
	if !checker.IsValid(username) {
		return
	}
	avail, err := checker.IsAvailable(username)
	if err != nil {
		log.Fatal(err)
	}
	if !avail {
		return
	}
	fmt.Printf("%q is valid and available on %s\n", username, checker)
}
