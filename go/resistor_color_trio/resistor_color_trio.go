package resistorcolortrio

import (
	"fmt"
	"math"
)

const (
	KILO_THRESHOLD = 1000
	MEGA_THRESHOLD = 1000000
	GIGA_THRESHOLD = 1000000000
)

func ColorValue(color string) int {
	colors := []string{
		"black",
		"brown",
		"red",
		"orange",
		"yellow",
		"green",
		"blue",
		"violet",
		"grey",
		"white",
	}

	for i, c := range colors {
		if color == c {
			return i
		}
	}

	return -1
}

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) (value int) {
	for i, color := range colors {
		colorValue := ColorValue(color)

		if i == 0 {
			colorValue *= 10
		}

		if i == 2 {
			value *= int(math.Pow(float64(10), float64(colorValue)))
			continue
		}

		value += colorValue
	}

	return value
}

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	resistorValue := Value(colors[:3])

	switch {
	case resistorValue < KILO_THRESHOLD:
		break
	case resistorValue%GIGA_THRESHOLD == 0:
		return fmt.Sprintf("%d gigaohms", resistorValue/GIGA_THRESHOLD)
	case resistorValue%MEGA_THRESHOLD == 0:
		return fmt.Sprintf("%d megaohms", resistorValue/MEGA_THRESHOLD)
	case resistorValue%KILO_THRESHOLD == 0:
		return fmt.Sprintf("%d kiloohms", resistorValue/KILO_THRESHOLD)
	}

	return fmt.Sprintf("%d ohms", resistorValue)
}
