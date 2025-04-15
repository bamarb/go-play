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

func Pow2Log(n uint) int {
	p := math.Ceil(math.Log2(float64(n)))
	return 1 << int(p)
}

func TestNearestPow2Log(t *testing.T) {
	assert.Equal(t, 4, NearestPow2Log(7))
	assert.Equal(t, 8, NearestPow2Log(11))
	assert.Equal(t, 8, NearestPow2Log(12))
	assert.Equal(t, 8, NearestPow2Log(15))
	assert.Equal(t, 1024, NearestPow2Log(1025))
}

func TestPow2Log(t *testing.T) {
	assert.Equal(t, 8, Pow2Log(7))
	assert.Equal(t, 16, Pow2Log(9))
	assert.Equal(t, 16, Pow2Log(12))
	assert.Equal(t, 16, Pow2Log(15))
	assert.Equal(t, 2048, Pow2Log(1025))
}

func TestNearestPow2Iter(t *testing.T) {
	assert.Equal(t, 4, NearestPow2Iter(7))
	assert.Equal(t, 8, NearestPow2Iter(11))
	assert.Equal(t, 8, NearestPow2Iter(12))
	assert.Equal(t, 8, NearestPow2Iter(15))
	assert.Equal(t, 1024, NearestPow2Iter(1025))
}

func TestMidPoint(t *testing.T) {
	lst := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sz := len(lst)
	mid := sz >> 1
	if mid != 5 {
		t.Errorf("got %d want %d", mid, 5)
	}
}

func Test2sComplement(t *testing.T) {
	for i := 0; i < 8; i++ {
		t.Logf("%08b", uint8(-i))
	}
}
