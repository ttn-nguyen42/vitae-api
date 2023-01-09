package about

import "time"

type GetResponse struct {
	Id         string `json:"id"`
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
	Birthday   time.Time `json:"birthday" time_format:"2006-01-02"`
	City       string    `json:"city" binding:"required"`
	Country    string    `json:"country" binding:"required"`
	Occupation string    `json:"occupation" binding:"required"`
}
