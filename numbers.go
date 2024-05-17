package spg

import (
	"math/rand"
	"strconv"
)

func (g Generator) randomNumber(dice int) int {
	return rand.Intn(dice) // #nosec G404: Intentianally using math/rand
}

// Gives back a strigified number
func (g Generator) StringNumber(max int) string {
	number := g.randomNumber(max)
	return strconv.Itoa(number)
}

// Gives back x chars in numbers
func (g Generator) StringNumberFixed(chars int) string {
	result := ""
	for i := 0; i < chars; i++ {
		result = result + g.StringNumber(10)
	}
	return result
}

func (g Generator) ExadecimalNumber() string {
	numbers := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	return numbers[g.randomNumber(16)]
}

func (g Generator) ExadecimalNumberFixed(chars int) string {
	result := ""
	for i := 0; i < chars; i++ {
		result = result + g.ExadecimalNumber()
	}
	return result
}
