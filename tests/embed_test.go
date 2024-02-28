package tests

import "testing"

type Fooer interface {
	Foo() string
}

type FooerImpl struct {
	str string
}

func (fi *FooerImpl) Foo() string {
	return fi.str
}

type Box struct {
	Fooer
	Baz int
}

func (b Box) GetBaz() int {
	return b.Baz
}

func TestInterfaceEmbedding(t *testing.T) {
	box := Box{Fooer: &FooerImpl{"Fi Fi"}, Baz: 420}
	t.Logf("Fooer says:%s\n", box.Foo())
}
