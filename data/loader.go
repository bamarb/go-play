package data

import "time"

type ChangeEventType string

const (
	CET_INSERT  ChangeEventType = "insert"
	CET_DELETE  ChangeEventType = "delete"
	CET_UPDATE  ChangeEventType = "update"
	CET_REPLACE ChangeEventType = "replace"
)

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
	Data D
	Type ChangeEventType
}

// A DataSource which provides items of type T
type DataSource[T any] interface {
	OnLoad(func([]T))
}
