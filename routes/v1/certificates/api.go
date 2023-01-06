package certificates

import (
	"Vitae/config/database"
	"Vitae/repositories/certificates"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	_ = certificates.New(database.Client)
	c.JSON(http.StatusOK, nil)
}