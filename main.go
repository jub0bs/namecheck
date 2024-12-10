package main

import (
	"fmt"
	"os"

	"github.com/jub0bs/namecheck/bluesky"
	"github.com/jub0bs/namecheck/github"
)

func main() {
	username := "jub0bs"
	valid := github.IsValid(username)
	fmt.Println("validity on GitHub:", valid)
	if valid {
		avail, err := github.IsAvailable(username)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(avail)
	}
	valid = bluesky.IsValid(username)
	fmt.Println("validity on Bluesky:", valid)
	if valid {
		avail, err := bluesky.IsAvailable(username)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(avail)
	}
}
