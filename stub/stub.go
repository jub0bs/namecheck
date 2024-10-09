package stub

import "net/http"

type SuccessfulGetter struct {
	StatusCode int
}

func (g *SuccessfulGetter) Get(_ string) (resp *http.Response, err error) {
	res := http.Response{
		StatusCode: g.StatusCode,
		Body:       http.NoBody,
	}
	return &res, nil
}
