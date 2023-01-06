package v1

// A generic response for sending a message as part of the response
type MessageResponse struct {
	Message string `json:"message" binding:"required"`
}

// A generic response for sending an ID as part of the response
type IdResponse struct {
	Id string `json:"id" binding:"required"`
}