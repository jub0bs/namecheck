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
		fmt.Fprintln(os.Stderr, "usage: namecheck <username>")
		os.Exit(1)
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
		available, err := gh.IsAvailable(username)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("availability of %q on GitHub: %t\n", username, available)
		}
	}
	var bs bluesky.Bluesky
	valid = bs.IsValid(username)
	fmt.Printf("validity of %q on Bluesky: %t\n", username, valid)
	if valid {
		available, err := bs.IsAvailable(username)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("availability of %q on Bluesky: %t\n", username, available)
		}
	}
}
