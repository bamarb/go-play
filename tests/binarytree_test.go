package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	bh1 = [...]int{1, 2, 3, 4, 5, 6, 7}
	bh2 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
)

func leaves(a []int) []int {
	leafStart := len(a) / 2
	return a[leafStart:]
}

func TestLeaves(t *testing.T) {
	l1 := leaves(bh1[:])
	l2 := leaves(bh2[:])
	require.ElementsMatch(t, []int{4, 5, 6, 7}, l1)
	require.ElementsMatch(t, []int{6, 7, 8, 9, 10}, l2)
}

func TestWeirdBitShift(t *testing.T) {
	var x uint64 = 2 << 10
	y := int(x << 1 >> 1)
	fmt.Printf("%d", y)
}
