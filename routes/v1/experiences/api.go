package experiences

import (
	"Vitae/config/database"
	"Vitae/repositories/experiences"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	_ = experiences.New(database.Client)
	c.JSON(http.StatusOK, nil)
}