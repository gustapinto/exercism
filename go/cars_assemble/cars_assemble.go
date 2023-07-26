package cars

const IndividualCarCost = 10000
const GroupOfTenCost = 95000

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate)) / 60
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	groupsCount := (carsCount / 10)
	groupsCost := groupsCount * GroupOfTenCost

	var individualCount int
	if groupsCount > 0 {
		individualCount = carsCount % (groupsCount * 10)
	} else {
		individualCount = carsCount
	}

	individualCost := individualCount * IndividualCarCost

	return uint(groupsCost + individualCost)
}
