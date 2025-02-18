package controller

import (
    "consulta_api_gov/services"
    "net/http"
    "github.com/gin-gonic/gin"
    "strconv"
)

func GarantiaSafraController(context *gin.Context) {
    cpf := context.DefaultQuery("cpf", "")

    paginaStr := context.DefaultQuery("pagina", "1")

    pagina, err := strconv.Atoi(paginaStr)
    if err != nil || pagina < 1 {
        context.JSON(http.StatusBadRequest, gin.H{"erro": "Número de página inválido"})
        return
    }

    if cpf == "" {
        context.JSON(http.StatusBadRequest, gin.H{"erro": "O CPF deve ser informado"})
        return
    }


    garantia, err := services.GetGarantiaSafra(cpf, paginaStr)
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
        return
    }

    context.HTML(http.StatusOK, "garantia.html", gin.H{
        "safraBeneficiario": garantia, 
    })
}