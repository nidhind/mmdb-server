package utils

import (
	"github.com/gin-gonic/gin"
	"mmdb-server/data_models"
	"mmdb-server/services"
	"log"
	"fmt"
)

func Authenticate(c *gin.Context) {
	credentials := data_models.LoginBody{}
	c.BindJSON(&credentials)

	query := data_models.LoginAuthQuery{
		UserId:credentials.UserId,
		Password:credentials.Password}

	var user data_models.UsersCollection

	session := services.GetNewSession()
	defer session.Close()

	err := session.DB(services.DB).C("users").Find(query).One(&user)
	if err != nil {
		log.Println(err)
		c.JSON(401, data_models.NewAPIError("INVALID_INPUT", "Invalid username or password"))
	} else {
		fmt.Println(user)
		c.JSON(200, user)
	}
}