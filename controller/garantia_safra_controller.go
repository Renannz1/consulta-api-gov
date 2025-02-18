package controller

import (
    "consulta_api_gov/services"
    "net/http"
    "github.com/gin-gonic/gin"
    "strconv"
)

// GarantiaSafraController processa a busca de dados do Garantia Safra
func GarantiaSafraController(context *gin.Context) {
    cpf := context.DefaultQuery("cpf", "")

    // Obtendo o número da página da consulta (com valor padrão 1)
    paginaStr := context.DefaultQuery("pagina", "1")

    // Verifica se o número da página é válido
    pagina, err := strconv.Atoi(paginaStr)
    if err != nil || pagina < 1 {
        context.JSON(http.StatusBadRequest, gin.H{"erro": "Número de página inválido"})
        return
    }

    // Verificando se o CPF está vazio
    if cpf == "" {
        context.JSON(http.StatusBadRequest, gin.H{"erro": "O CPF deve ser informado"})
        return
    }

    // Passa o CPF e o número da página para o serviço
    garantia, err := services.GetGarantiaSafra(cpf, paginaStr)
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
        return
    }

    // Exibindo os dados encontrados
    context.HTML(http.StatusOK, "garantia.html", gin.H{
        "safraBeneficiario": garantia, // Passa a lista de resultados
    })
}