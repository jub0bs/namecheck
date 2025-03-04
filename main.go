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
		os.Exit(1)
	}
	username := os.Args[1]
	valid := github.IsValid(username)
	fmt.Printf("validity of %q GitHub: %t\n", username, valid)
	if valid {
		avail, err := github.IsAvailable(username)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			fmt.Printf("available of %q GitHub: %t\n", username, avail)
		}
	}
	valid = bluesky.IsValid(username)
	fmt.Printf("validity of %q Bluesky: %t\n", username, valid)
	if valid {
		avail, err := bluesky.IsAvailable(username)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			fmt.Printf("available of %q Bluesky: %t\n", username, avail)
		}
	}
}
