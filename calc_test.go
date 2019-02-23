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
