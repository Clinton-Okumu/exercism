package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	successfulCarsMade := float64(productionRate) * float64(successRate/100)
	return float64(successfulCarsMade)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	carsPerMinute := (CalculateWorkingCarsPerHour((productionRate), successRate) / 60)
	return int(carsPerMinute)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	tens := carsCount / 10
	ones := carsCount % 10
	totalCost := (tens * 95000) + (ones * 10000)
	return uint(totalCost)
}
