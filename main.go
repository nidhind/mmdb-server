package main

import (
	"github.com/gin-gonic/gin"
	"mmdb-server/services"
)

func main() {

	api := gin.Default()

	// Establish a MongoDB session
	services.InitMongo();

	// Mount API routes
	MountRoutes(api);

	// Listen and Server in 0.0.0.0:8080
	api.Run(":8080")
}
