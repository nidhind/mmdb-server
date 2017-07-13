package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mmdb-server/data_models"
	"mmdb-server/services"
)

var db = services.Session

func Authenticate(c *gin.Context) {
	credentials := data_models.LoginBody{}
	c.BindJSON(&credentials)

	query := struct {

	}{}
	err := data_models.NewAPIError("INVALID_INPUT", "Invalid user name or password")
	c.JSON(403, err)
}