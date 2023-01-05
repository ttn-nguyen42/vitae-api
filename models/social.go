package models

type SocialMedia struct {
	Platform     string `json:"platform" binding:"required"`
	User         string `json:"user" binding:"required"`
	Url          Url    `json:"url" binding:"required"`
	ColoredIcon  Url    `json:"coloredIcon"`
	MonotoneIcon Url    `json:"monotoneIcon"`
}
