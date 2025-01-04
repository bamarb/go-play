package ifx

type Named interface {
	name() string
}

type User struct {
	fName string
}

func (u *User) name() string {
	return u.fName
}

func getName(x Named) string {
	return x.name()
}

type Car struct {
	make string
}

func (c *Car) name() string {
	return c.make
}

var _ Named = (*User)(nil)
