package main

import (
	"github.com/gin-gonic/gin"

	"mmdb-server/utils"
	"fmt"
	"github.com/gin-contrib/cors"
)

func MountRoutes(app *gin.Engine) {

	// Enable CROS
	app.Use(cors.Default())

	// Ping test
	app.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	authorised := app.Group("/")
	authorised.Use(func(c *gin.Context) {
		c.Next()
	})
	authorised.POST("admin", utils.Authenticate)
	authorised.Use(func(c *gin.Context) {
		if len(c.Errors) != 0 {
			fmt.Println(c.Errors)
		} else {
			c.Next()
		}
	})

	// Handle 404
	app.Use(func(c *gin.Context) {
		c.String(404, "No middleware responded!")
	})

}