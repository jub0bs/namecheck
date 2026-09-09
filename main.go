package main

import (
	"fmt"
	"log"

	"github.com/jub0bs/namecheck/bluesky"
	"github.com/jub0bs/namecheck/github"
)

func main() {
	username := "jub0bs"
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
