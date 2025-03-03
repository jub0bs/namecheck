package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func main() {
	username := "jub0bs-"
	re, err := regexp.Compile("^[A-Za-z0-9-]{3,39}$")
	if err != nil {
		log.Fatal(err)
	}
	if strings.HasPrefix(username, "-") ||
		strings.HasSuffix(username, "-") ||
		strings.Contains(username, "--") ||
		!re.MatchString(username) {
		return
	}
	fmt.Println(username)
}
