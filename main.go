package main

import (
	"fmt"
	"os"

	"github.com/jub0bs/namecheck/github"
	"github.com/jub0bs/namecheck/reddit"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: namecheck <username>")
		os.Exit(1)
	}
	username := os.Args[1]
	valid := github.IsValid(username)
	fmt.Printf("%q is valid on GitHub: %t\n", username, valid)
	if valid {
		avail, err := github.IsAvailable(username)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%q is available on GitHub: %t\n", username, avail)
		}
	}
	valid = reddit.IsValid(username)
	fmt.Printf("%q is valid on Reddit: %t\n", username, valid)
	if valid {
		avail, err := reddit.IsAvailable(username)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%q is available on Reddit: %t\n", username, avail)
		}
	}
}
