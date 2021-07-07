package namecheck

import "fmt"

type Validator interface {
	IsValid(username string) bool
}

type Availabler interface {
	IsAvailable(username string) (bool, error)
}

type Checker interface {
	Validator
	Availabler
	fmt.Stringer
}

type ErrUnknownAvailability struct {
	Username string
	Platform string
	Cause    error
}

func (err *ErrUnknownAvailability) Error() string {
	const tmpl = "unknown availability of %q on %s: %v"
	return fmt.Sprintf(tmpl, err.Username, err.Platform, err.Cause)
}
