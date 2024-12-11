package stub

import "net/http"

type SuccessfulClient struct {
	StatusCode int
}

func (sc *SuccessfulClient) Get(string) (*http.Response, error) {
	res := http.Response{
		StatusCode: sc.StatusCode,
		Body:       http.NoBody,
	}
	return &res, nil
}
