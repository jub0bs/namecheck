package main

import (
	"fmt"

	"github.com/jub0bs/namecheck/github"
)

func main() {
	fmt.Println(github.IsValid("jub0bs"))
	fmt.Println(github.IsValid("jub  0bs"))
	fmt.Println(github.IsValid("jub___0bs"))
}
