package tests

import (
	"testing"
	"tests/pkga"
	"tests/pkgb"
)

func TestStructurallySimilarStructs(t *testing.T) {
	sa := pkga.Sizea{Width: 100, Height: 100}
	sb := pkgb.Sizeb{Width: 100, Height: 100}
	if sa == pkga.Sizea(sb) {
		t.Logf("Structs are same \n")
	}
}
