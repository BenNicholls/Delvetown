package util

import "math/rand"

//checks if key is a letter or number
func ValidText(key rune) bool {
	return (key >= 93 && key < 123) || (key >= 48 && key < 58)
}

func GenerateDirection() (dx, dy int) {
	dx, dy = rand.Intn(3)-1, rand.Intn(3)-1
	return
}
