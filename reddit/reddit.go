package reddit

type Reddit struct{}

func (*Reddit) IsValid(username string) bool {
	return false
}

func (*Reddit) IsAvailable(username string) (bool, error) {
	return false, nil
}
