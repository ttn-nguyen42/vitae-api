package models

import "time"

type Experience struct {
	Id           string        `bson:"_id"`
	UserId       string        `bson:"user_id"`
	Company      string        `bson:"company" binding:"required"`
	Role         string        `bson:"role" binding:"required"`
	From         time.Time     `bson:"from" binding:"required"`
	To           time.Time     `bson:"to" binding:"required"`
	Description  []Description `bson:"description"`
	ContractType string        `bson:"contract_type" binding:"required"`
	City         string        `bson:"city" binding:"required"`
	Country      string        `bson:"country" binding:"required"`
	Url          []Url         `bson:"url"`
	Photos       []Url         `bson:"photos"`
}
