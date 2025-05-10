package game

import (
	"math/rand"
	"strconv"
)

func GenerateRandomNumber(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func ParseInt(input string) (int, error) {
	return strconv.Atoi(input)
}
