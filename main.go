package main

import (
	"consulta_api_gov/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	routes.SetupRoutes(r)

	r.Run(":8080")
}
