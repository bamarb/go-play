package tests

import (
	"fmt"
	"testing"
)

type MyInt int

func (m MyInt) String() string {
	return "42"
}

func PrintInt[T ~int](val T) {
	fmt.Printf("Type: %T value: %d\n", val, val)
}

func TestUnd(t *testing.T) {
	var mi MyInt = 1
	PrintInt(1)
	PrintInt(mi)
}
