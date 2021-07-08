package stub

import (
	"io"
	"net/http"
	"os"
)

type SuccessfulClient struct {
	StatusCode int
}

func (sc *SuccessfulClient) Get(url string) (*http.Response, error) {
	res := http.Response{
		Body:       io.NopCloser(os.Stdin),
		StatusCode: sc.StatusCode,
	}
	return &res, nil

}
