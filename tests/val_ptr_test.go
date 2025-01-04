package tests

import "testing"

type PassByVal struct {
	m   map[string]int
	str string
	s   []string
	v   int
}

func NewPBV() PassByVal {
	return PassByVal{v: 1, m: map[string]int{}, s: []string{"jingo"}, str: "inito"}
}

func PByVal(v PassByVal) {
	// Let's mutate shit here
	v.m["amar"] = 42
	v.s[0] = "singo"
	v.str = "ringo"
}

func TestPByVal(t *testing.T) {
	pbv := NewPBV()
	t.Logf("Before: %v\n", pbv)
	PByVal(pbv)
	t.Logf("After: %v\n", pbv)
}
