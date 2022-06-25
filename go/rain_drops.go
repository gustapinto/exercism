package raindrops

import "fmt"

func Convert(number int) string {
	sound := ""

	if number%3 == 0 {
		sound += "Pling"
	}
	if number%5 == 0 {
		sound += "Plang"
	}
	if number%7 == 0 {
		sound += "Plong"
	}

	if sound == "" {
		// If the sound is empty then the number doesnt have 3, 5 or 7 as factors
		return fmt.Sprintf("%d", number)
	}

	return sound
}
