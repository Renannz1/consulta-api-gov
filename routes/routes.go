package routes

import (
	"consulta_api_gov/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// pagian inicial
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// pessoa fisica
	router.GET("/form-pessoa-fisica", func(c *gin.Context) {
		c.HTML(200, "form_pessoa_fisica.html", nil)
	})
	router.GET("/pessoa-fisica", controller.PessoaFisicaController)

	// beneficiario auxilio emergencial
	router.GET("/form-beneficiario-auxilio", func(c *gin.Context) {
		c.HTML(200, "form_beneficiario_auxilio.html", nil)
	})
	router.GET("/beneficiario-auxilio", controller.AuxilioEmergencialController)

}
