package about

import (
	"Vitae/config/database"
	"Vitae/repositories/about"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	_ = about.New(database.Client)
	c.JSON(http.StatusOK, nil)
}
