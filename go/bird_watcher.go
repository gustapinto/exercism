package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0
	for _, t := range birdsPerDay {
		total += t
	}

	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	endDay := week * 7
	startDay := endDay - 7

	total := 0
	for i := startDay; i < endDay; i++ {
		total += birdsPerDay[i]
	}

	return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	fixedBirdsPerDay := birdsPerDay
	for i, t := range birdsPerDay {
		if !(i%2 == 0) {
			continue
		}

		fixedBirdsPerDay[i] = t + 1
	}

	return fixedBirdsPerDay
}
