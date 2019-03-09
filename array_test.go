package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxInts(t *testing.T) {
	input := []int{1, 0, 100, 5, 10, -2, -100, 100}
	actual := MaxInts(input...)

	expected := 100
	assert.Equal(t, expected, actual)
}

func TestMaxIntsWithIdx(t *testing.T) {
	input := []int{1, 0, 100, 5, 10, -2, -100, 100}
	actual := struct {
		Value int
		Index int
	}{}
	actual.Value, actual.Index = MaxIntsWithIdx(input...)

	expected := struct {
		Value int
		Index int
	}{100, 2}
	assert.Equal(t, expected, actual)
}

func TestMinInts(t *testing.T) {
	input := []int{1, 0, 100, 5, 10, -2, -100, 100}
	actual := MinInts(input...)

	expected := -100
	assert.Equal(t, expected, actual)
}

func TestMinIntsWithIdx(t *testing.T) {
	input := []int{1, 0, 100, 5, 10, -2, -100, 100}
	actual := struct {
		Value int
		Index int
	}{}
	actual.Value, actual.Index = MinIntsWithIdx(input...)

	expected := struct {
		Value int
		Index int
	}{-100, 6}
	assert.Equal(t, expected, actual)
}

func TestSumInts(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	actual := SumInts(input)

	expected := 11 * 10 / 2
	assert.Equal(t, expected, actual)
}

func TestSumInts2d(t *testing.T) {
	input := [][]int{
		[]int{1, 2, 3, 4, 5},
		[]int{6, 7, 8, 9, 10},
	}
	actual := SumInts2d(input)

	expected := 55
	assert.Equal(t, expected, actual)
}

func TestReverseInts(t *testing.T) {
	input := []int{1, 0, 100, 5, 10, -2, -100, 100}
	actual := ReverseInts(input)

	expected := []int{100, -100, -2, 10, 5, 100, 0, 1}
	assert.Equal(t, expected, actual)
}

func TestExtendInts(t *testing.T) {
	arr1, arr2 := []int{1, 1, 1}, []int{2, 2, 2}
	actual := ExtendInts(arr1, arr2)

	expected := []int{1, 1, 1, 2, 2, 2}
	assert.Equal(t, expected, actual)
}

func TestInts2d(t *testing.T) {
	actual := Ints2d(20, 10)
	assert.Equal(t, 20, len(actual))
	assert.Equal(t, 10, len(actual[0]))
}

func TestStrInts2d(t *testing.T) {
	mat := [][]int{
		[]int{1, 2, 3, 4, 5},
		[]int{6, 7, 8, 9, 10},
	}
	actual := StrInts2d(mat)
	expected := "[1 2 3 4 5]\n[6 7 8 9 10]"
	assert.Equal(t, expected, actual)
}

func TestMemsetInts1d(t *testing.T) {
	nums := make([]int, 10)
	MemsetInts1d(nums, 5)
	for _, n := range nums {
		assert.Equal(t, 5, n)
	}
}

func TestMemsetInts2d(t *testing.T) {
	nums := Ints2d(10, 10)
	MemsetInts2d(nums, 5)
	for _, row := range nums {
		for _, col := range row {
			assert.Equal(t, 5, col)
		}
	}
}

func TestContainsInt(t *testing.T) {
	// exist
	nums, num := []int{1, 4, 6, 7, 2}, 4
	actual := ContainsInt(nums, num)
	expected := true
	assert.Equal(t, expected, actual)

	nums, num = []int{1, 4, 6, 2}, 2
	actual = ContainsInt(nums, num)
	expected = true
	assert.Equal(t, expected, actual)

	// not exist
	nums, num = []int{1, 4, 6, 7, 2}, 100
	actual = ContainsInt(nums, num)
	expected = false
	assert.Equal(t, expected, actual)

	nums, num = []int{6, 2}, -1
	actual = ContainsInt(nums, num)
	expected = false
	assert.Equal(t, expected, actual)
}
