package main

import (
	"fmt"

	"github.com/jub0bs/namecheck/github"
)

func main() {
	username := "jub0bs"
	if github.IsValid(username) {
		fmt.Println(username)
	}
}
