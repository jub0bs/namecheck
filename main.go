package main

import (
	"fmt"

	"github.com/jub0bs/namecheck/bluesky"
	"github.com/jub0bs/namecheck/github"
)

func main() {
	username := "jub0bs"
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
