package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jub0bs/namecheck/bluesky"
	"github.com/jub0bs/namecheck/github"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, "usage: namecheck <username>\n")
		os.Exit(1)
	}
	username := os.Args[1]
	gh := github.GitHub{
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
	if !gh.IsValid(username) {
		return
	}
	avail, err := gh.IsAvailable(username)
	if err != nil {
		log.Fatal(err)
	}
	if !avail {
		return
	}
	fmt.Printf("%q is valid and available on GitHub\n", username)
	var bs bluesky.Bluesky
	if !bs.IsValid(username) {
		return
	}
	avail, err = bs.IsAvailable(username)
	if err != nil {
		log.Fatal(err)
	}
	if !avail {
		return
	}
	fmt.Printf("%q is valid and available on Bluesky\n", username)
}
