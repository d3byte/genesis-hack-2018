package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID        primitive.ObjectID `json:"-" bson:"_id,omitempty" example:"bson id"`
	Email     string             `json:"email" example:"test@test.ru"`
	Role      int                `json:"-" example:"0" default:"0"`
	Password  string             `json:"password,omitempty" example:"12345"`
	Token     string             `json:"-"`
	Confirmed *bool              `json:"-"`
	Time      time.Time          `json:"-"`
}

type UserWithHiddenFields struct {
	ID        primitive.ObjectID `json:"-" bson:"_id,omitempty" example:"bson id"`
	Email     string             `json:"email" example:"test@test.ru"`
	Role      int                `json:"role,omitempty" example:"0" default:"0"`
	Token     string             `json:"-"`
	Confirmed *bool              `json:"-"`
	Time      time.Time          `json:"-"`
}

type UserLogin struct {
	Email    string `json:"email,omitempty" example:"test@test.ru" binding:"required"`
	Password string `json:"password,omitempty" example:"12345" binding:"required"`
}
