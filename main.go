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
	if valid, _ := regexp.MatchString("^[a-zA-Z0-9-]{3,39}$", username); !valid {
		return
	}
	fmt.Println(username)
}
