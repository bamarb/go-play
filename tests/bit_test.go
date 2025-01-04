package tests

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsPowerOfTwo(value uint) bool {
	return value&(value-1) == 0
}

func NearestPow2Log(n uint) int {
	p := math.Floor(math.Log2(float64(n)))
	return int(math.Pow(2, p))
}

func NearestPow2Iter(n uint) int {
	res := 0
	if n < 1 {
		return 0
	}
	for i := n; i >= 1; i-- {
		// if i is divisible by 2 return it
		if IsPowerOfTwo(i) {
			res = int(i)
			break
		}
	}
	return res
}

func TestNearestPow2Log(t *testing.T) {
	assert.Equal(t, 4, NearestPow2Log(7))
	assert.Equal(t, 8, NearestPow2Log(11))
	assert.Equal(t, 8, NearestPow2Log(12))
	assert.Equal(t, 8, NearestPow2Log(15))
	assert.Equal(t, 1024, NearestPow2Log(1025))
}

func TestNearestPow2Iter(t *testing.T) {
	assert.Equal(t, 4, NearestPow2Iter(7))
	assert.Equal(t, 8, NearestPow2Iter(11))
	assert.Equal(t, 8, NearestPow2Iter(12))
	assert.Equal(t, 8, NearestPow2Iter(15))
	assert.Equal(t, 1024, NearestPow2Iter(1025))
}

func TestMidPoint(t *testing.T) {
	lst := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sz := len(lst)
	mid := sz >> 1
	if mid != 4 {
		t.Errorf("got %d want %d", mid, 4)
	}
}
