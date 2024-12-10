package main

import (
	"fmt"
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
	fmt.Println("validity on GitHub:", valid)
	if valid {
		avail, err := gh.IsAvailable(username)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(avail)
	}
	var bs bluesky.Bluesky
	valid = bs.IsValid(username)
	fmt.Println("validity on Bluesky:", valid)
	if valid {
		avail, err := bs.IsAvailable(username)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(avail)
	}
}
