package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jub0bs/namecheck/bluesky"
	"github.com/jub0bs/namecheck/github"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, "usage: namecheck <username> [more-usernames]\n")
		os.Exit(1)
	}
	usernames := make(map[string]bool)
	for _, name := range os.Args[1:] {
		usernames[name] = true
	}
	for username := range usernames {
		if !github.IsValid(username) {
			continue
		}
		avail, err := github.IsAvailable(username)
		if err != nil {
			log.Fatal(err)
		}
		if !avail {
			continue
		}
		fmt.Printf("%q is valid and available on GitHub\n", username)
		if !bluesky.IsValid(username) {
			continue
		}
		avail, err = bluesky.IsAvailable(username)
		if err != nil {
			log.Fatal(err)
		}
		if !avail {
			continue
		}
		fmt.Printf("%q is valid and available on Bluesky\n", username)
	}
}
