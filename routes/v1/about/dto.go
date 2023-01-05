package about

import "time"

type GetResponse struct {
	FirstName  string `json:"firstName" binding:"required"`
	LastName   string `json:"lastName" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	City       string `json:"city" binding:"required"`
	Country    string `json:"country" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
}

type PostRequest struct {
	FirstName  string    `json:"firstName" binding:"required"`
	LastName   string    `json:"lastName" binding:"required"`
	Email      string    `json:"email" binding:"required,email"`
	Birthday   time.Time `json:"birthday"`
	City       string    `json:"city" binding:"required"`
	Country    string    `json:"country" binding:"required"`
	Occupation string    `json:"occupation" binding:"required"`
}