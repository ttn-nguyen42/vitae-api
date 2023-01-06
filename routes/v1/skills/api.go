package skills

import (
	"Vitae/config/database"
	"Vitae/repositories/skills"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	_ = skills.New(database.Client)
	c.JSON(http.StatusOK, nil)
}