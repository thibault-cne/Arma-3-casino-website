package rouletteservices

import (
	"fmt"
	"math/rand"
)

func calculateRouletteColor() string {
	result := randInt(0, 37)
	color := ""

	if result == 0 {
		color = "green"
	} else if result%2 == 0 {
		color = "red"
	} else {
		color = "black"
	}

	fmt.Printf("Rand number : %d\n", result)

	return color
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
