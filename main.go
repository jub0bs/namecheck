package main

import (
	"fmt"
	"log"

	"github.com/jub0bs/namecheck/bluesky"
	"github.com/jub0bs/namecheck/github"
)

func main() {
	username := "jub0bs"
	valid := github.IsValid(username)
	fmt.Printf("validity of %q on GitHub: %t\n", username, valid)
	if valid {
		avail, err := github.IsAvailable(username)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("availability of %q on GitHub: %t\n", username, avail)
	}
	valid = bluesky.IsValid(username)
	fmt.Printf("validity of %q on Bluesky: %t\n", username, valid)
	if valid {
		avail, err := bluesky.IsAvailable(username)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("availability of %q on Bluesky: %t\n", username, avail)
	}
}
