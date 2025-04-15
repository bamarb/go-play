package tests

import (
	"cmp"
	"slices"
	"testing"
)

var intVals = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

func BinSearch[T ~[]E, E cmp.Ordered](s T, target E) (int, bool) {
	return slices.BinarySearch(s, target)
}

func TestCmp(t *testing.T) {
	a, b := 1, 1
	if cmp.Less(a, b) {
		t.Log("a < b")
	} else if cmp.Less(b, a) {
		t.Log("b < a")
	} else {
		t.Log("a == b")
	}
}

func TestOverflow(t *testing.T) {
	i, j := 12, 20
	z := int(uint(i+j) >> 1)
	t.Logf("i: %d,j: %d, z : %d\n", i, j, z)
	z = (i + j) >> 1
	t.Logf("i: %d,j: %d, z : %d\n", i, j, z)
	i, j = 1024, 1024
}
