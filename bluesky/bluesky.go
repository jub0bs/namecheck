package bluesky

type Bluesky struct{}

func (*Bluesky) IsValid(username string) bool { return false }

func (*Bluesky) IsAvailable(username string) (bool, error) { return false, nil }
