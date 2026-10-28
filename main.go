package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	username := "jub0bs"
	fmt.Println(IsValid(username))
}

func IsValid(username string) (bool, error) {
	if strings.HasPrefix(username, "-") ||
		strings.HasSuffix(username, "-") ||
		strings.Contains(username, "--") {
		return false, nil
	}
	ok, err := regexp.MatchString("^[A-Za-z0-9-]{3,39}$", username)
	if err != nil {
		return false, err
		//log.Fatal("impossible to determine validity of username")
	}
	if !ok {
		return false, nil
	}
	return true, nil
}
