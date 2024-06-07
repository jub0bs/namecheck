package main

import (
	"fmt"
	"os"

	"github.com/jub0bs/namecheck/github"
	"github.com/jub0bs/namecheck/reddit"
)

func main() {
	const username = "jub0bs"
	if !github.IsValid(username) {
		fmt.Printf("%q is not valid on GitHub\n", username)
	} else {
		fmt.Printf("%q is valid on GitHub\n", username)
		avail, err := github.IsAvailable(username)
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

	if !reddit.IsValid(username) {
		fmt.Printf("%q is not valid on Reddit\n", username)
	} else {
		fmt.Printf("%q is valid on Reddit\n", username)
		avail, err := reddit.IsAvailable(username)
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
