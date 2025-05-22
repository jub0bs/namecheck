package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jub0bs/namecheck/bluesky"
	"github.com/jub0bs/namecheck/github"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <username>\n", os.Args[0])
		os.Exit(1)
	}
	username := os.Args[1]
	gh := github.GitHub{
		Client: http.DefaultClient,
	}
	valid := gh.IsValid(username)
	fmt.Printf("validity of %q on GitHub: %t\n", username, valid)
	if valid {
		avail, err := gh.IsAvailable(username)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("availability of %q on GitHub: %t\n", username, avail)
	}
	var bs bluesky.Bluesky
	valid = bs.IsValid(username)
	fmt.Printf("validity of %q on Bluesky: %t\n", username, valid)
	if valid {
		avail, err := bs.IsAvailable(username)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("availability of %q on Bluesky: %t\n", username, avail)
	}
}
