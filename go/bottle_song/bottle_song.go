package bottlesong

import (
	"fmt"
	"strings"
)

var (
	numbers = map[int]string{
		10: "Ten",
		9: "Nine",
		8: "Eight",
		7: "Seven",
		6: "Six",
		5: "Five",
		4: "Four",
		3: "Three",
		2: "Two",
		1: "One",
		0: "No",
	}
)

func Recite(startBottles, takeDown int) []string {
	verses := make([]string, 0)
	sb := startBottles

	for i := 1; i <= takeDown; i++ {
		bottle := "bottle"
		if sb > 1 {
			bottle = "bottles"
		}
	
		verses = append(verses, fmt.Sprintf("%s green %s hanging on the wall,", numbers[sb], bottle))
		verses = append(verses, fmt.Sprintf("%s green %s hanging on the wall,", numbers[sb], bottle))
		verses = append(verses, fmt.Sprintf("And if one green bottle should accidentally fall,"))
	
		sb -= 1
		if sb == 1 {
			bottle = "bottle"
		} else {
			bottle = "bottles"
		}
	
		verses = append(verses, fmt.Sprintf("There'll be %s green %s hanging on the wall.", strings.ToLower(numbers[sb]), bottle))

		if takeDown > 1 && i != takeDown {
			verses = append(verses, "")
		}
	}

	return verses
}
