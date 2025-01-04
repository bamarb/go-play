package reflectx

import (
	"fmt"
	"reflect"
	"testing"
)

func outln(x any, t ...*testing.T) {
	if len(t) == 0 {
		fmt.Printf("type is '%T', value: %v\n", x, x)
		return
	}
	t[0].Logf("type is '%T', value: %v\n", x, x)
}

func ptr[T any](v T) *T {
	return &v
}

func TestReflectionTypes(t *testing.T) {
	outln([5]byte{}, t) // Type: [5]uint8
	outln("binga", t)
	outln(ptr("binga"), t)
	outln(nil, t)
	outln(struct{ x string }{}, t)
	var x any = 3.1415
	typ := reflect.TypeOf(x)
	val := reflect.ValueOf(x)
	t.Logf("Type of any x: %s value-kind: %s\n", typ, val.Kind())
	x = "binga"
	typ = reflect.TypeOf(x)
	val = reflect.ValueOf(x)
	t.Logf("Type of any x: %s value-kind: %s\n", typ, val.Kind())
}
