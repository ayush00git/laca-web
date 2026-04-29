package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Course struct {
	ID			primitive.ObjectID		`bson:"_id,omitempty" json:"_id"`
	Code		string					`bson:"code" json:"code" binding:"required"`
	MaxSeats	uint16					`bson:"maxSeats" json:"maxSeats" binding:"required"`
	Students	[]primitive.ObjectID	`bson:"students" json:"students"`
	CreatedAt	time.Time				`bson:"created_at" json:"created_at"`
}
