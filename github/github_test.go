package github_test

import (
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/jub0bs/namecheck/github"
)

func TestIsValid(t *testing.T) {
	type TestCase struct {
		//desc     string
		username string
		want     bool
	}
	testCases := map[string]TestCase{
		"contains two consecutive hyphens": {"jub0bs--on-GitHub", false},
		"starts with a hyphen":             {"-jub0bs", false},
		"ends with a hyphen":               {"jub0bs-", false},
		"too short":                        {"ab", false},
		"too long":                         {strings.Repeat("a", 40), false},
		"contains illegal chars":           {"jub&bs", false},
		"all good":                         {"jub0bs", true},
	}
	var gh github.GitHub
	for desc, tc := range testCases {
		f := func(t *testing.T) {
			got := gh.IsValid(tc.username)
			if got != tc.want {
				const tmpl = "github.IsValid(%q): got %t; want %t"
				t.Errorf(tmpl, tc.username, got, tc.want)
			}
		}
		t.Run(desc, f)
	}
}

type StubClient struct {
	StatusCode int
	Err        error
}

func (sc *StubClient) Do(req *http.Request) (*http.Response, error) {
	if sc.Err != nil {
		return nil, sc.Err
	}
	res := http.Response{
		StatusCode: sc.StatusCode,
		Body:       http.NoBody, // try to comment this out and see
	}
	return &res, nil
}

func TestIsAvailableErrorCase(t *testing.T) {
	gh := github.GitHub{
		Client: &StubClient{
			Err: errors.New("oh no"),
		},
	}
	avail, err := gh.IsAvailable("whatever")
	if err == nil || avail {
		t.Errorf("got %t, %v; want false, some non-nil error", avail, err)
	}
}

func TestIsAvailable404(t *testing.T) {
	gh := github.GitHub{
		Client: &StubClient{
			StatusCode: http.StatusNotFound,
		},
	}
	avail, err := gh.IsAvailable("whatever")
	if err != nil || !avail {
		t.Errorf("got %t, %v; want true, nil", avail, err)
	}
}

func TestIsAvailable200(t *testing.T) {
	gh := github.GitHub{
		Client: &StubClient{
			StatusCode: http.StatusOK,
		},
	}
	avail, err := gh.IsAvailable("whatever")
	if err != nil || avail {
		t.Errorf("got %t, %v; want false, nil", avail, err)
	}
}

func TestIsAvailableOtherStatusCode(t *testing.T) {
	gh := github.GitHub{
		Client: &StubClient{
			StatusCode: 299,
		},
	}
	avail, err := gh.IsAvailable("whatever")
	if err == nil || avail {
		t.Errorf("got %t, %v; want false, some non-nil error", avail, err)
	}
}
