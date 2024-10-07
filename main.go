package main

import (
	"fmt"

	"github.com/jub0bs/namecheck/github"
)

func main() {
	fmt.Println(github.IsValid("jub0bs")) // true
	fmt.Println(github.IsValid("-_-"))    // false
}
