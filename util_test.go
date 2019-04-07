package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllComb(t *testing.T) {
	n := 2
	actual := AllComb(n)

	expected := [][]int{
		[]int{0, 0},
		[]int{0, 1},
		[]int{1, 0},
		[]int{1, 1},
	}

	assert.ElementsMatch(t, expected, actual)
}

func TestAllCombChan(t *testing.T) {
	n := 2
	nPatterns := 1 << uint(n)

	// prepare channels
	c := make(chan []int, nPatterns)
	go AllCombChan(n, c)

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

func TestPermuteInts(t *testing.T) {
	nums := []int{1, 2, 3}
	actual := PermuteInts(nums)
	expected := [][]int{
		[]int{1, 2, 3},
		[]int{1, 3, 2},
		[]int{2, 1, 3},
		[]int{2, 3, 1},
		[]int{3, 1, 2},
		[]int{3, 2, 1},
	}

	assert.ElementsMatch(t, expected, actual)
}
