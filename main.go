package main

import (
	"fmt"

	"github.com/jub0bs/namecheck/github"
	"github.com/jub0bs/namecheck/reddit"
)

func main() {
	username := "jub0bs"
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
