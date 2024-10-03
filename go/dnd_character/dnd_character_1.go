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
	var throws = [4]int{}

	for i := 0; i < 4; i++ {
		throws[i] = (rand.IntN(D6_SIDES) + 1)
	}

	lowestIndex := 0
	for i := range throws {
		if throws[i] < throws[lowestIndex] {
			lowestIndex = i
		}
	}

	sum := 0
	for i, t := range throws {
		if i == lowestIndex {
			continue
		}

		sum += t
	}
	
	return sum
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
