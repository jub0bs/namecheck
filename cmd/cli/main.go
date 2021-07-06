package main

import (
	"fmt"

	"github.com/jub0bs/namecheck/github"
	"github.com/jub0bs/namecheck/twitter"
)

func main() {
	username := "jub0bs"
	fmt.Println(twitter.IsValid(username))
	fmt.Println(github.IsValid(username))
}
