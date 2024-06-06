package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	username := "jub0bs"
	if !IsValid(username) {
		return
	}
	fmt.Println(username)
}

func IsValid(username string) bool {
	if strings.HasPrefix(username, "-") ||
		strings.HasSuffix(username, "-") ||
		strings.Contains(username, "--") {
		return false
	}
	if valid, _ := regexp.MatchString("^[a-zA-Z0-9-]{3,39}$", username); !valid {
		return false
	}
	return true
}
