package main

import (
	"fmt"
	"os"

	"github.com/jub0bs/namecheck/bluesky"
	"github.com/jub0bs/namecheck/github"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: namecheck <username>")
		os.Exit(1)
	}
	username := os.Args[1]
	valid := github.IsValid(username)
	fmt.Printf("validity of %q on GitHub: %t\n", username, valid)
	if valid {
		available, err := github.IsAvailable(username)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("availability of %q on GitHub: %t\n", username, available)
		}
	}
	valid = bluesky.IsValid(username)
	fmt.Printf("validity of %q on Bluesky: %t\n", username, valid)
	if valid {
		available, err := bluesky.IsAvailable(username)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("availability of %q on Bluesky: %t\n", username, available)
		}
	}
}
