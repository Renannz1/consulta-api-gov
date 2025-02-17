package controller

import (
	"consulta_api_gov/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetExpensesController lida com a requisição para exibir os gastos públicos
func GetExpensesController(c *gin.Context) {
	expenses, err := services.GetExpenses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar despesas"})
		return
	}

	c.JSON(http.StatusOK, expenses)
}
