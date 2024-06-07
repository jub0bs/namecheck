package reddit

import "net/http"

type Reddit struct {
	Client *http.Client
}

func (*Reddit) IsValid(username string) bool {
	return false
}

func (*Reddit) IsAvailable(username string) (bool, error) {
	return false, nil
}

func (*Reddit) String() string {
	return "Reddit"
}
