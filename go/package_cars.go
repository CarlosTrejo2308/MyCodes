package cars

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	rate := successRate(speed)
	return float64(speed) * 221.0 * rate
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	hour := CalculateProductionRatePerHour(speed)
	return int(hour) / 60
}

// successRate is used to calculate the ratio of an item being created without
// error for a given speed
func successRate(speed int) float64 {
	switch {
	case speed == 0:
		return 0
	case speed <= 4:
		return 1
	case speed <= 8:
		return .9
	default:
		return .77
	}
}
