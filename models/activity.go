package models

import "time"

type Activity struct {
	Role        string        `json:"role" binding:"required"`
	Organizer   string        `json:"organizer" binding:"required"`
	From        time.Time     `json:"from" binding:"required"`
	To          time.Time     `json:"to" binding:"required"`
	Description []Description `json:"description"`
	Url         []Url         `json:"url"`
}
