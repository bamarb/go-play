package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Base struct {
	// Oid Object Id field set by mongo when this entity is persisted
	Oid     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Created time.Time          `json:"created,omitempty" bson:"created,omitempty"`
	Updated time.Time          `json:"updated,omitempty" bson:"updated,omitempty"`
}

func NewBase() Base {
	return Base{
		Oid:     primitive.NewObjectID(),
		Created: time.Now(),
		Updated: time.Now(),
	}
}
