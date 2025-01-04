package ifx

import (
	"fmt"
	"io"
	"os"
	"testing"
)

// Interface Assertions
// i.(T) where i is a var holding interface type and T is a type you are asserting

func TestAssertPrimitive(t *testing.T) {
	var x any = "bleh"
	if v, ok := x.(string); ok {
		t.Logf("value stored: %s", v)
	} else {
		t.Fail()
	}

	var y io.Reader = os.Stdout
	if v, ok := y.(io.Writer); ok {
		fmt.Fprintln(v, "yay")
	}

	if v, ok := y.(*os.File); ok {
		t.Logf("Got a File with Name %s\n", v.Name())
	}

	// Dynamic Interface assertion

	if v, ok := y.(interface{ Write([]byte) (int, error) }); ok {
		t.Logf(" val staisfies Writer interface , v%\n", v)
	}
}

func TestAssertComposite(t *testing.T) {
	type Junk struct {
		x int
		y int
	}

	type Junks []Junk

	var x any = Junk{}
	if v, ok := x.(Junk); ok {
		t.Logf("junk value %v\n", v)
	}

	var y any = Junks{Junk{1, 1}, Junk{2, 2}}
	if junks, ok := y.(Junks); ok {
		t.Logf("Junks %v\n", junks)
	}
}
