package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	username := "jub0bs"
	if strings.Contains(username, "--") ||
		strings.HasPrefix(username, "-") ||
		strings.HasSuffix(username, "-") {
		return
	}
	matched, _ := regexp.MatchString("^[a-zA-Z0-9-]{3,39}$", username)
	if !matched {
		return
	}
	fmt.Println(username)
}
