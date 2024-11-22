package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string             `bson:"title" json:"title"`
	Body      string             `bson:"body" json:"body" `
	Author    string             `bson:"author" json:"author"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type AddBlog struct {
	Title  string `bson:"title" json:"title" validate:"required"`
	Body   string `bson:"body" json:"body" validate:"required"`
	Author string `bson:"author" json:"author" validate:"required"`
}
