package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	username := "jub0bs"
	if strings.HasPrefix(username, "-") ||
		strings.HasSuffix(username, "-") ||
		strings.Contains(username, "--") {
		return
	}
	re := regexp.MustCompile("^[a-zA-Z0-9-]{3,39}$")
	if !re.MatchString(username) {
		return
	}
	fmt.Println(username)
}
