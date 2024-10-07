package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	const username = "jub0bs"
	if strings.HasPrefix(username, "-") ||
		strings.HasSuffix(username, "-") ||
		strings.Contains(username, "--") {
		return
	}
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9-]{3,39}$", username); !ok {
		return
	}
	fmt.Println(username)
}
