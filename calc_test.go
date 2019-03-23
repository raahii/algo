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

func TestGcd(t *testing.T) {
	actual := Gcd(20, 32)
	expected := 4
	assert.Equal(t, expected, actual)

	values := []int{290021904, 927964716, 826824516, 817140688}

	z := values[0]
	for i := 1; i < len(values); i++ {
		z = Gcd(values[i], z)
	}
	actual = z
	expected = 92
	assert.Equal(t, expected, actual)
}
