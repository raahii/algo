package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPow(t *testing.T) {
	actual := Pow(2, 4)
	expected := 16
	assert.Equal(t, expected, actual)
}

func TestAbs(t *testing.T) {
	actual := Abs(-100)
	expected := 100
	assert.Equal(t, expected, actual)

	actual = Abs(255)
	expected = 255
	assert.Equal(t, expected, actual)
}
