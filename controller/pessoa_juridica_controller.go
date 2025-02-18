package controller

import (
    "consulta_api_gov/services"
    "net/http"
    "github.com/gin-gonic/gin"
)

func PessoaJuridicaController(context *gin.Context) {
    cnpj := context.DefaultQuery("cnpj", "")


    if cnpj == "" {
        context.JSON(http.StatusBadRequest, gin.H{"erro": "O CNPJ deve ser informado"})
        return
    }

    pessoaJuridica, err := services.GetPessoaJuridica(cnpj)
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
        return
    }

    context.HTML(http.StatusOK, "pessoajuridica.html", gin.H{
        "pessoaJuridica": pessoaJuridica,
    })
}