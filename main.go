package main

import (
	"fmt"
	"regexp"
	"strings"
)

var re = regexp.MustCompile("^[a-zA-Z0-9-]{3,39}$")

func main() {
	username := "jub0bs"
	fmt.Println(IsValid(username))
}

func IsValid(username string) bool {
	return !strings.Contains(username, "--") &&
		!strings.HasPrefix(username, "-") &&
		!strings.HasSuffix(username, "-") &&
		re.MatchString(username)
}
