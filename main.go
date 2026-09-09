package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

const looksValid = "^[A-Za-z0-9-]{3,39}$"

func main() {
	fmt.Println(IsValid("jub0bs"))
	fmt.Println(IsValid("jub0bs-"))
	fmt.Println(IsValid("-jub0bs"))
	fmt.Println(IsValid("jub--0bs"))
	fmt.Println(IsValid("ju"))
	fmt.Println(IsValid(strings.Repeat("a", 40)))
	fmt.Println(IsValid("jub*bs"))
}

func IsValid(username string) bool {
	if strings.HasPrefix(username, "-") ||
		strings.Contains(username, "--") ||
		strings.HasSuffix(username, "-") {
		return false
	}
	match, err := regexp.MatchString(looksValid, username)
	if err != nil {
		log.Fatal(err)
	}
	return match
}
