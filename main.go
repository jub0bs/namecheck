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
		fmt.Fprint(os.Stderr, "usage: namecheck <username>\n")
		os.Exit(1)
	}
	username := os.Args[1]
	if !github.IsValid(username) {
		return
	}
	avail, err := github.IsAvailable(username)
	if err != nil {
		log.Fatal(err)
	}
	if !avail {
		return
	}
	fmt.Printf("%q is valid and available on GitHub\n", username)
	if !bluesky.IsValid(username) {
		return
	}
	avail, err = bluesky.IsAvailable(username)
	if err != nil {
		log.Fatal(err)
	}
	if !avail {
		return
	}
	fmt.Printf("%q is valid and available on Bluesky\n", username)
}
