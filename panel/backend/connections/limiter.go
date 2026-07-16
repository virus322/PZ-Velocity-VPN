package connections

func IsAllowed(current int, max int) bool {
	if max <= 0 {
		return true
	}

	return current < max
}
