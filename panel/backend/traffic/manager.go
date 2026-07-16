package traffic

func CalculateTotal(upload int64, download int64) int64 {
	return upload + download
}

func IsLimitReached(total int64, limit int64) bool {
	return total >= limit
}
