package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

const looksValid = "^[A-Za-z0-9-]{3,39}$"

func main() {
	username := "jub0bs"
	if strings.HasPrefix(username, "-") ||
		strings.Contains(username, "--") ||
		strings.HasSuffix(username, "-") {
		return
	}
	match, err := regexp.MatchString(looksValid, username)
	if err != nil {
		log.Fatal(err)
	}
	if !match {
		return
	}
	fmt.Println(username)
}
