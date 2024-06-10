package namecheck

import (
	"fmt"
	"net/http"
)

type Getter interface {
	Get(url string) (resp *http.Response, err error)
}

type Validator interface {
	IsValid(string) bool
}

type Availabler interface {
	IsAvailable(string) (bool, error)
}

type Checker interface {
	Validator
	Availabler
	fmt.Stringer
}
