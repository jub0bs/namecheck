package namecheck

import "net/http"

type Getter interface {
	Do(*http.Request) (resp *http.Response, err error)
}
