package main

import (
	"fmt"
	"net/http"
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
	gh := github.GitHub{
		Client: http.DefaultClient,
	}
	valid := gh.IsValid(username)
	fmt.Printf("%q is valid on GitHub: %t\n", username, valid)
	if valid {
		avail, err := gh.IsAvailable(username)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%q is available on GitHub: %t\n", username, avail)
		}
	}
	var re reddit.Reddit
	valid = re.IsValid(username)
	fmt.Printf("%q is valid on Reddit: %t\n", username, valid)
	if valid {
		avail, err := re.IsAvailable(username)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%q is available on Reddit: %t\n", username, avail)
		}
	}
}
