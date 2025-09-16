package bluesky

import "context"

type Bluesky struct{}

func (*Bluesky) IsValid(username string) bool { return false }

func (*Bluesky) IsAvailable(ctxt context.Context, username string) (bool, error) { return false, nil }

func (gh *Bluesky) String() string {
	return "Bluesky"
}
