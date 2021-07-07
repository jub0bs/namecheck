package namecheck

import "fmt"

type ErrUnknownAvailability struct {
	Username string
	Platform string
	Cause    error
}

func (err *ErrUnknownAvailability) Error() string {
	const tmpl = "unknown availability of %q on %s: %v"
	return fmt.Sprintf(tmpl, err.Username, err.Platform, err.Cause)
}
