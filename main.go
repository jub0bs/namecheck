package main

import (
	"fmt"
	"strings"

	"github.com/jub0bs/namecheck/github"
)

func main() {
	fmt.Println(github.IsValid("jub0bs"))
	fmt.Println(github.IsValid("jub0bs-"))
	fmt.Println(github.IsValid("-jub0bs"))
	fmt.Println(github.IsValid("jub--0bs"))
	fmt.Println(github.IsValid("ju"))
	fmt.Println(github.IsValid(strings.Repeat("a", 40)))
	fmt.Println(github.IsValid("jub*bs"))
}
