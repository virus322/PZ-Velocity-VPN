package expire

import "time"

func IsExpired(expireDate string) bool {
	if expireDate == "" {
		return false
	}

	t, err := time.Parse("2006-01-02", expireDate)
	if err != nil {
		return false
	}

	return time.Now().After(t)
}
