package algo

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadIntF(t *testing.T) {
	f := strings.NewReader("3\n")

	actual := ReadIntF(f)

	expected := 3
	assert.Equal(t, expected, actual)
}

func TestReadIntsF(t *testing.T) {
	f := strings.NewReader("3 5 7\n")

	actual := ReadIntsF(f, 3)

	expected := []int{3, 5, 7}
	assert.Equal(t, expected, actual)
}

func TestReadLineF(t *testing.T) {
	str := "010101001001001000001010101010011111111111101100101000010101001000000000\n"
	f := strings.NewReader(str)

	expected := "010101001001001000001010101010011111111111101100101000010101001000000000"
	actual := ReadLineF(f)
	assert.Equal(t, expected, actual)
}

func TestReadLinesF(t *testing.T) {
	str := "オッス！おら悟空！\nいっちょやってみっか！\n"
	f := strings.NewReader(str)

	expected := []string{
		"オッス！おら悟空！",
		"いっちょやってみっか！",
	}
	actual := ReadLinesF(f, 2)
	assert.Equal(t, expected, actual)
}
