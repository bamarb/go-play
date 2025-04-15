package tests

import (
	"fmt"
	"testing"
	"unsafe"
)

// Tests for functions  in slices package

func baseAddr[E any](s []E) uintptr {
	return uintptr(unsafe.Pointer(&s[0]))
}

func lastAddr[E any](s []E) uintptr {
	return uintptr(unsafe.Pointer(&s[len(s)-1]))
}

func TestBaseAddr(t *testing.T) {
	bs := []byte{1, 2, 3}
	t.Logf("baddr:%p  baddr:%#x", &bs[0], baseAddr(bs))
}

func unsafeOverlaps[S any](a, b []S) bool {
	if len(a) == 0 || len(b) == 0 {
		return false
	}
	esz := unsafe.Sizeof(a[0])
	if esz == 0 {
		return false
	}
	baseAddra := baseAddr(a)
	lastAddra := lastAddr(a)
	baseAddrb := baseAddr(b)
	lastAddrb := lastAddr(b)
	fmt.Printf("ba:%#x la:%#x\n", baseAddra, lastAddra)
	fmt.Printf("bb:%#x lb:%#x\n", baseAddrb, lastAddrb)
	ret := baseAddra <= lastAddrb && baseAddrb <= lastAddra
	fmt.Printf("ret:%t\n", ret)
	return ret
}

func TestUnsafeOverlap(t *testing.T) {
	x := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	if !unsafeOverlaps(x, x) {
		t.Error("Expecting true for x x")
	}
	y := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	if unsafeOverlaps(x, y) {
		t.Error("Expecting false for x y")
	}

	z := append(x[:0:0], 1, 2, 3) // allocates a new backing array
	if unsafeOverlaps(x, z) {
		t.Error("Expecting false (no overlap) for x z")
	}

	z = x[2:3]
	if !unsafeOverlaps(x, z) {
		t.Error("Expecting overlap (true) for x and subslice z")
	}

	z = x[1:1] //n:n is a zero len subslice
	if unsafeOverlaps(x, z) {
		t.Logf("Wasnt expecting an overlap zero len subslice")
	}
}
