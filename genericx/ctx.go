package genericx

// Adding Context Keys and Values

// Key is a generic key type associated with a  specific value type
type Key[V any] struct{ name *string }

func NewKey[V any](name string) Key[V] {
	return Key[V]{&name}
}
