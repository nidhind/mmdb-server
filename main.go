package main

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/gin-contrib/cors"
)

type LoginBody struct {
	Access_token string `form:"user" json:"user" binding:"required"`
}

func main() {

	api := gin.Default()

	// Enable CROS
	api.Use(cors.Default())

	// Example for binding JSON ({"user": "manu", "password": "123"})

	// Ping test
	api.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	authorized := api.Group("/")

	authorized.Use(func(c *gin.Context) {
		c.Next()
	})
	authorized.POST("admin", func(c *gin.Context) {
		var json LoginBody
		if c.BindJSON(&json) == nil {
			if json.Access_token == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		}
	})

	api.Use(func(c *gin.Context) {
		c.String(404, "No middleware responded!")
	})
	// Listen and Server in 0.0.0.0:8080
	api.Run(":8080")
}
