package github_test

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/jub0bs/namecheck/github"
)

func ExampleGitHub_IsValid() {
	var gh github.GitHub
	fmt.Println(gh.IsValid("jub0bs"))
	// Output: true
}

func TestIsValid(t *testing.T) {
	type TestCase struct {
		username string
		want     bool
	}
	testCases := map[string]TestCase{
		"contains two consecutive hyphens": {"jub0bs--on-GitHub", false},
		"too short":                        {"ab", false},
		"too long":                         {strings.Repeat("a", 40), false},
		"contains illegal chars":           {"^^^", false},
		"starts with a hyphen":             {"-jub0bs", false},
		"ends with a hyphen":               {"jub0bs-", false},
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

func (c *StubClient) Do(req *http.Request) (*http.Response, error) {
	if c.Err != nil {
		return nil, c.Err
	}
	res := http.Response{
		StatusCode: c.StatusCode,
		Body:       http.NoBody,
	}
	return &res, nil
}

func TestIsAvailableError(t *testing.T) {
	gh := github.GitHub{
		Client: &StubClient{Err: errors.New("oh no")},
	}
	avail, err := gh.IsAvailable("whatever")
	if err == nil || avail {
		t.Errorf("got %t, %s; want false, some non-nil error", avail, err)
	}
}

func TestIsAvailable404(t *testing.T) {
	gh := github.GitHub{
		Client: &StubClient{StatusCode: http.StatusNotFound},
	}
	avail, err := gh.IsAvailable("whatever")
	if err != nil || !avail {
		t.Errorf("got %t, %s; want false, some non-nil error", avail, err)
	}
}

func TestIsAvailable200(t *testing.T) {
	gh := github.GitHub{
		Client: &StubClient{StatusCode: http.StatusOK},
	}
	avail, err := gh.IsAvailable("whatever")
	if err != nil || avail {
		t.Errorf("got %t, %s; want false, some non-nil error", avail, err)
	}
}

func TestIsAvailableUnexpectedStatusCode(t *testing.T) {
	gh := github.GitHub{
		Client: &StubClient{StatusCode: http.StatusBadGateway},
	}
	avail, err := gh.IsAvailable("whatever")
	if err == nil || avail {
		t.Errorf("got %t, %s; want false, some non-nil error", avail, err)
	}
}
