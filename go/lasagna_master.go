package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, layerTime int) int {
	if layerTime == 0 {
		layerTime = 2
	}

	return len(layers) * layerTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}

	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) (scaledAmounts []float64) {
	for _, amount := range amounts {
		scaledAmount := amount * (float64(portions) / 2.0)
		scaledAmounts = append(scaledAmounts, scaledAmount)
	}

	return
}
