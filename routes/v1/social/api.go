package social

import (
	"Vitae/config/database"
	"Vitae/repositories/social"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	_ = social.New(database.Client)
	c.JSON(http.StatusOK, nil)
}