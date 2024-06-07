package reddit

import (
	"github.com/jub0bs/namecheck"
)

type Reddit struct {
	Client namecheck.Getter
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
