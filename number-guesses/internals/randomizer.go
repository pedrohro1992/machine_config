package internals

import (
	"math/rand"
	"time"
)

// RandomizerNumbers returns a array with a random value
func RandomizerNumbers(s []int) []int {
	result := make([]int, len(s))

	copy(result, s)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})

	return result
}
