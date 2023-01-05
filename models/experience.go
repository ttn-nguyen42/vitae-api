package models

import "time"

type Experience struct {
	Company      string        `json:"company" binding:"required"`
	Role         string        `json:"role" binding:"required"`
	From         time.Time     `json:"from" binding:"required"`
	To           time.Time     `json:"to" binding:"required"`
	Description  []Description `json:"description"`
	ContractType string        `json:"contractType" binding:"required"`
	City         string        `json:"city" binding:"required"`
	Country      string        `json:"country" binding:"required"`
	Url          []Url         `json:"url"`
	Photos       []Url         `json:"photos"`
}
