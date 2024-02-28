package randx

import (
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
)

// UUID returns a new UUID or panics
func UUID() string {
	return uuid.New().String()
}

func ULID() string {
	return ulid.Make().String()
}
