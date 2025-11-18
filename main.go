package main

import (
	"fmt"
	"os"

	"github.com/jub0bs/namecheck/bluesky"
	"github.com/jub0bs/namecheck/github"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <username>\n", os.Args[0])
		os.Exit(2)
	}
	usernames := make(map[string]struct{})
	for _, u := range os.Args[1:] {
		usernames[u] = struct{}{}
	}
	for username := range usernames {
		valid := github.IsValid(username)
		fmt.Printf("validity of %q on GitHub: %t\n", username, valid)
		if valid {
			avail, err := github.IsAvailable(username)
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
}
