package main

import (
	"consulta_api_gov/config"
	"consulta_api_gov/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
