package main

import (
	"fmt"
	"regexp"
	"strings"
)

var re = regexp.MustCompile("^[A-Za-z0-9-]{3,39}$")

func main() {
	username := "jub0bs-"
	if !IsValid(username) {
		return
	}
	fmt.Println(username)
}

func IsValid(username string) bool {
	return !strings.HasPrefix(username, "-") &&
		!strings.HasSuffix(username, "-") &&
		!strings.Contains(username, "--") &&
		re.MatchString(username)
}
