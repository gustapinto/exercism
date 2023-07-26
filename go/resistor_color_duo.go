package resistorcolorduo

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
	for i, color := range colors[:2] {
		colorValue := ColorValue(color)

		if i == 0 {
			colorValue *= 10
		}

		value += colorValue
	}

	return value
}
