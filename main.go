package main

import (
	"fmt"
	"log"
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
	ok, err := regexp.MatchString("^[A-Za-z0-9-]{3,39}$", username)
	if err != nil {
		log.Fatal("impossible to determine validity of username")
	}
	if !ok {
		return
	}
	fmt.Println(username)
}
