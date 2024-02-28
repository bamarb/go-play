package tests

import (
	"fmt"
	"testing"
)

type Container struct {
	i int
	s string
}

func (c Container) byValMethod() {
	fmt.Printf("byValMethod got &c=%p, &(c.s)=%p\n", &c, &(c.s))
}

func (c *Container) byPtrMethod() {
	fmt.Printf("byPtrMethod got &c=%p, &(c.s)=%p\n", c, &(c.s))
}

func TestValueVsPointerReceiver(t *testing.T) {
	var c Container
	fmt.Printf("in main &c=%p, &(c.s)=%p\n", &c, &(c.s))
	c.byValMethod()
	c.byPtrMethod()
}
