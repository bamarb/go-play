package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func add(a, b int) int {
	fmt.Println("add")
	return a + b
}

func scale(a, b int) (int, int) {
	fmt.Println("scale")
	return a * 2, b * 2
}

func TestMulAdd(t *testing.T) {
	i := add(scale(2, 2))
	assert.Equal(t, 8, i)
}
