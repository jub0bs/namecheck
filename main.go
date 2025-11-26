package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/jub0bs/namecheck/bluesky"
	"github.com/jub0bs/namecheck/github"
)

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
	valid := gh.IsValid(username)
	fmt.Printf("validity of %q on GitHub: %t\n", username, valid)
	if valid {
		avail, err := gh.IsAvailable(username)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("availability of %q on GitHub: %t\n", username, avail)
		}
	}
	valid = bluesky.IsValid(username)
	fmt.Printf("validity of %q on Bluesky: %t\n", username, valid)
	if valid {
		avail, err := bluesky.IsAvailable(username)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("availability of %q on Bluesky: %t\n", username, avail)
		}
	}
}
