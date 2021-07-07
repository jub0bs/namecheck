package twitter

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strings"

	"github.com/jub0bs/namecheck"
)

type Twitter struct{}

var re = regexp.MustCompile("^[0-9A-Z_a-z]{4,15}$")

func containsNoIllegalPattern(username string) bool {
	return !strings.Contains(strings.ToLower(username), "twitter")
}

func looksGood(username string) bool {
	return re.MatchString(username)
}

func (*Twitter) IsValid(username string) bool {
	return looksGood(username) && containsNoIllegalPattern(username)
}

func (tw *Twitter) IsAvailable(username string) (bool, error) {
	endpoint := "https://europe-west6-namechecker-api.cloudfunctions.net/userlookup?username=" + username
	resp, err := http.Get(endpoint)
	if err != nil {
		err1 := namecheck.ErrUnknownAvailability{
			Username: username,
			Platform: tw.String(),
			Cause:    err,
		}
		return false, &err1
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err1 := namecheck.ErrUnknownAvailability{
			Username: username,
			Platform: tw.String(),
			Cause:    errors.New("unexpected response from API"),
		}
		return false, &err1
	}
	var dto struct {
		Data interface{} `json:"data"`
	}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&dto); err != nil {
		err1 := namecheck.ErrUnknownAvailability{
			Username: username,
			Platform: tw.String(),
			Cause:    err,
		}
		return false, &err1
	}
	// the absence of a data field in the response body indicates the username's availability
	return dto.Data == nil, nil
}

func (*Twitter) String() string {
	return "Twitter"
}
