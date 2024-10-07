package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(IsValid("jub0bs")) // true
	fmt.Println(IsValid("-_-"))    // false
}

func IsValid(username string) bool {
	if strings.HasPrefix(username, "-") ||
		strings.HasSuffix(username, "-") ||
		strings.Contains(username, "--") {
		return false
	}
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9-]{3,39}$", username); !ok {
		return false
	}
	return true
}
