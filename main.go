package main

import (
	"fmt"
	"regexp"
	"strings"
)

var re = regexp.MustCompile("^[a-zA-Z0-9-]{3,39}$")

func main() {
	fmt.Println(IsValid("jub0bs"))
	fmt.Println(IsValid("jub  0bs"))
	fmt.Println(IsValid("jub___0bs"))
}

func IsValid(username string) bool {
	if strings.HasPrefix(username, "-") ||
		strings.HasSuffix(username, "-") ||
		strings.Contains(username, "--") {
		return false
	}
	return re.MatchString(username)
}
