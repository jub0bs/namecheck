package namecheck

import "net/http"

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}
