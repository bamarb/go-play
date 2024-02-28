package tests

import (
	"cmp"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

// multi sort
type Bid struct {
	From     string
	Price    float64
	Priority int
}

var bids1 = []Bid{
	{"A", 2.01, 2}, {"B", 2.01, 2}, {"C", 2.00, 10}, {"D", 1.99, 5}, {"E", 3.99, 6},
}

var bids2 = []Bid{
	{"A", 2.01, 2}, {"B", 2.00, 2}, {"C", 2.00, 3}, {"D", 1.99, 5}, {"E", 1.99, 6},
}

func merge(slcA []Bid, slcB []Bid) []Bid {
	merged := []Bid{}
	merged = append(merged, slcA...)
	merged = append(merged, slcB...)
	return merged
}

func BidComparatorByPrio(i, j Bid) int {
	if n := cmp.Compare(i.Priority, j.Priority); n != 0 {
		return n
	}
	// if Priorities are equal sort by price
	return cmp.Compare(i.Price, j.Price)
}

func TestSort_multi(t *testing.T) {
	t.Logf("unsorted:%v\n", bids1)
	slices.SortFunc(bids1, BidComparatorByPrio)
	t.Logf("sorted:%v\n", bids1)
}

func TestSliceLenAndCap(t *testing.T) {
	strSlc := make([]string, 0, 3)
	strSlc = append(strSlc, "a")
	strSlc = append(strSlc, "b")
	strSlc = append(strSlc, "c")
	t.Logf("%v", strSlc)
}

func TestAppend(t *testing.T) {
	ret := merge(nil, nil)
	require.Empty(t, ret)
	ret = merge(nil, bids2)
	require.Len(t, ret, len(bids2))
}
