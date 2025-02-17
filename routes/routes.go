package routes

import (
	"consulta_api_gov/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.GET("/form-pessoa-fisica", func(c *gin.Context) {
		c.HTML(200, "form_pessoa_fisica.html", nil)
	})
	router.GET("/pessoa-fisica", controller.PessoaFisicaController)
}
