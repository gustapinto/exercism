package dndcharacter

import (
	"math"
	"math/rand/v2"
)

const D6_SIDES = 6

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	mod := (float64(score) - 10) / 2
	
	return int(math.Floor(mod))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	sum := 0
	lowestValue := D6_SIDES // Init as max to avoid never updating

	for i := 0; i < 4; i++ {
		value := (rand.IntN(D6_SIDES) + 1)

		if value < lowestValue {
			lowestValue = value
		}

		sum += value
	}
	
	return sum - lowestValue
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	c := Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		Hitpoints:    0,
	}

	c.Hitpoints = 10 + Modifier(c.Constitution)

	return c
}
