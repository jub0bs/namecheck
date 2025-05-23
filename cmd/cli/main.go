package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/jub0bs/namecheck/bluesky"
	"github.com/jub0bs/namecheck/github"
)

type Checker interface {
	IsValid(string) bool
	IsAvailable(string) (bool, error)
	fmt.Stringer
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <username>\n", os.Args[0])
		os.Exit(1)
	}
	username := os.Args[1]
	gh := github.GitHub{Client: http.DefaultClient}
	bs := bluesky.Bluesky{}
	var checkers []Checker
	for range 20 {
		checkers = append(checkers, &gh, &bs)
	}
	var wg sync.WaitGroup
	for _, checker := range checkers {
		wg.Add(1)
		go check(checker, username, &wg)
	}
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
		log.Fatal(err)
	}
	fmt.Printf("availability of %q on %s: %t\n", username, checker, avail)
}
