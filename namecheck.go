package namecheck

import "net/http"

type Client interface {
	Get(url string) (resp *http.Response, err error)
}
