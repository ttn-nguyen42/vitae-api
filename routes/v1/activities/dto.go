package activities

import (
	"Vitae/models"
	"time"
)

type GetResponse struct {
	Id          string   `json:"_id,omitempty" copier:"-"`
	Role        string               `json:"role" binding:"required"`
	Organizer   string               `json:"organizer" binding:"required"`
	From        time.Time            `json:"from" binding:"required"`
	To          time.Time            `json:"to" binding:"required"`
	Description []models.Description `json:"description"`
	Url         []models.Url         `json:"url"`
}

type PostRequest struct {
	Role        string               `json:"role" binding:"required"`
	Organizer   string               `json:"organizer" binding:"required"`
	From        time.Time            `json:"from" binding:"required"`
	To          time.Time            `json:"to" binding:"required"`
	Description []models.Description `json:"description"`
	Url         []models.Url         `json:"url"`
}
