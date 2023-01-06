package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Experience struct {
	Id           primitive.ObjectID `bson:"_id,omitempty" copier:"-"`
	UserId       primitive.ObjectID `bson:"user_id"`
	Company      string             `bson:"company" binding:"required"`
	Role         string             `bson:"role" binding:"required"`
	From         time.Time          `bson:"from" binding:"required"`
	To           time.Time          `bson:"to" binding:"required"`
	Description  []Description      `bson:"description"`
	ContractType string             `bson:"contract_type" binding:"required"`
	City         string             `bson:"city" binding:"required"`
	Country      string             `bson:"country" binding:"required"`
	Url          []Url              `bson:"url"`
	Photos       []Url              `bson:"photos"`
}
