package tests

import "testing"

type Point struct {
	x int
	y int
}

type Heap[T any] struct {
	data []T
	size int
}

func TestNilSlice(t *testing.T) {
	h := &Heap[int]{}
	t.Logf("nil slice elem access : %d", h.data[0])
}

func MutateSlice(s []Point, t *testing.T) {
	for i := 0; i < 10; i++ {
		// t.Logf("appending point to slice %d\n", i)
		s[i] = Point{i, i + 1}
	}
}

func TestMutateSlice(t *testing.T) {
	t.SkipNow()
	slc := []Point{}
	t.Logf("Len before mutate %d\n", len(slc))
	MutateSlice(slc, t)
	t.Logf("Len after mutate %d\n", len(slc))
	if len(slc) != 0 {
		t.FailNow()
	}
}

func TestMutateSlice_alloc(t *testing.T) {
	t.SkipNow()
	slc := make([]Point, 0, 10)
	t.Logf("Len before mutate %d\n", len(slc))
	MutateSlice(slc, t)
	t.Logf("Len after mutate %d, %v\n", len(slc), slc)
	if len(slc) != 10 {
		t.FailNow()
	}
}

func TestSliceBasic(t *testing.T) {
	slc := make([]Point, 4)
	t.Logf("Len:%d Cap:%d", len(slc), cap(slc))

	for i := 0; i < 2; i++ {
		// slc = append(slc, Point{i, i + 1})
		slc[i] = Point{i, i + 1}
	}
	t.Logf("Len after mutate %d, %v\n", len(slc), slc)
	slc = make([]Point, 0, 3)
	for i := 0; i < 4; i++ {
		slc = append(slc, Point{i, i + 1})
	}
	t.Logf("Len after mutate %d, %d  %v\n", len(slc), cap(slc), slc)

	for i := 0; i < len(slc); i++ {
		t.Logf("%+v", slc[i])
	}
}
