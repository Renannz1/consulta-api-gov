package routes

import (
	"consulta_api_gov/controller"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configura as rotas da aplicação
func SetupRoutes(router *gin.Engine) {
	router.GET("/despesas", controller.GetExpensesController)
	router.GET("/pessoa/:cpf", controller.PessoaFisicaController)
}
