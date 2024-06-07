package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jub0bs/namecheck/github"
	"github.com/jub0bs/namecheck/reddit"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: namecheck <username>")
		os.Exit(1)
	}
	username := os.Args[1]
	gh := github.GitHub{
		Client: http.DefaultClient,
	}
	if !gh.IsValid(username) {
		fmt.Printf("%q is not valid on GitHub\n", username)
	} else {
		fmt.Printf("%q is valid on GitHub\n", username)
		avail, err := gh.IsAvailable(username)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unknown availibility of %q on GitHub", username)
			return
		}
		if !avail {
			fmt.Printf("%q is not available on GitHub\n", username)
		} else {
			fmt.Printf("%q is available on GitHub\n", username)
		}
	}
	re := reddit.Reddit{
		Client: http.DefaultClient,
	}
	if !re.IsValid(username) {
		fmt.Printf("%q is not valid on Reddit\n", username)
	} else {
		fmt.Printf("%q is valid on Reddit\n", username)
		avail, err := re.IsAvailable(username)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unknown availibility of %q on Reddit", username)
			return
		}
		if !avail {
			fmt.Printf("%q is not available on Reddit\n", username)
		} else {
			fmt.Printf("%q is available on Reddit\n", username)
		}
	}
}
