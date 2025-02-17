package controller

import (
	"consulta_api_gov/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PessoaFisicaController(context *gin.Context) {
	cpf := context.DefaultQuery("cpf", "")

	if cpf == "" {
		context.JSON(http.StatusBadRequest, gin.H{"erro": "O cpf não pode ser vazio"})
	}

	pessoa, err := services.GetPessoaFisica(cpf)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	context.HTML(http.StatusOK, "pessoafisica.html", gin.H{
		"pessoa": pessoa,
	})
}
