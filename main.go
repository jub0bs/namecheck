package main

import (
	"fmt"
	"log"
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
	valid, err := regexp.MatchString("^[A-Z0-9a-z-]{3,39}$", username)
	if err != nil {
		log.Fatal(err)
	}
	return valid
}
