package syncx

import (
	"testing"

	"github.com/puzpuzpuz/xsync/v3"
)

type Point struct {
	x int
	y int
}

func TestConcStrMap(t *testing.T) {
	m := xsync.NewMapOf[Point, int]()
	m.Store(Point{42, 42}, 42)
}
