package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinComb(t *testing.T) {
	n := 2
	nPatterns := 1 << uint(n)

	// prepare channels
	c := make(chan []int, nPatterns)
	go BinComb(n, c)

	actual := Ints2d(nPatterns, n)
	i := 0
	for p := range c {
		actual[i] = p
		i++
	}

	expected := [][]int{
		[]int{0, 0},
		[]int{0, 1},
		[]int{1, 0},
		[]int{1, 1},
	}

	assert.ElementsMatch(t, expected, actual)
}
