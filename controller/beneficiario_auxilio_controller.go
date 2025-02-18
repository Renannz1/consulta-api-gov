package controller

import (
	"consulta_api_gov/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuxilioEmergencialController(context *gin.Context) {
	codigoBeneficiario := context.DefaultQuery("codigoBeneficiario", "")

	if codigoBeneficiario == "" {
		context.JSON(http.StatusBadRequest, gin.H{"erro": "O código do beneficiário não pode ser vazio"})
		return
	}

	pagina := context.DefaultQuery("pagina", "1")
	if pagina == "" {
		pagina = "1"
	}

	dados, err := services.GetAuxilioEmergencialPorCPF(codigoBeneficiario, pagina)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	context.HTML(http.StatusOK, "beneficiario_auxilio.html", gin.H{
		"dados": dados,
	})
}
