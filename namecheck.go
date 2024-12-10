package namecheck

import "net/http"

type Getter interface {
	Get(string) (*http.Response, error)
}
