package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Experience struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	Experience string             `bson:"experience" json:"experience" `
	CreatedAt  time.Time          `bson:"createdAt" json:"-"`
	UpdatedAt  time.Time          `bson:"updatedAt" json:"-"`
}

type ExperienceDto struct {
	Experience string    `json:"experience"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}


