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
		return
	}

	pessoa, _ := services.GetPessoaFisica(cpf)

	context.HTML(http.StatusOK, "pessoafisica.html", gin.H{
		"pessoa": pessoa,
	})
}