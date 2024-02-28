package data

import "time"

type ChangeEventType string

const (
	CET_INSERT  ChangeEventType = "insert"
	CET_DELETE  ChangeEventType = "delete"
	CET_UPDATE  ChangeEventType = "update"
	CET_REPLACE ChangeEventType = "replace"
)

type Loader[T any] interface {
	// OnLoad a consumer of the Loaded type
	OnLoad(func(T any))
}

type Watcher[E any] interface {
	Watch() <-chan ChangeEvent[E]
}

// Marker interface
type Event interface{}

type TypedEvent[T any, D any] interface {
	Type() T
	Data() D
	Ts() time.Time
}

type ChangeEvent[D any] struct {
	Type ChangeEventType
	Data D
}
