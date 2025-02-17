package controller

import (
	"consulta_api_gov/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PessoaFisicaController(context *gin.Context) {
	cpf := context.Param("cpf")

	pessoa, err := services.GetPessoaFisica(cpf)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	context.JSON(http.StatusOK, pessoa)
}
