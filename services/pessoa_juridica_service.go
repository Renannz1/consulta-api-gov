package services

import (
    "encoding/json"
    "fmt"
    "net/http"
    "consulta_api_gov/config"
    "consulta_api_gov/models"
)

func GetPessoaJuridica(cnpj string) (*models.PessoaJuridica, error) {
    config := config.LoadConfig()
    url := fmt.Sprintf("%s/api-de-dados/pessoa-juridica?cnpj=%s", config.PortalAPIURL, cnpj)

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, fmt.Errorf("erro ao criar requisição: %v", err)
    }

    // Adiciona a chave da API no cabeçalho
    req.Header.Add("chave-api-dados", config.PortalAPIKey)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("erro ao fazer requisição: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("erro na requisição, status code: %d", resp.StatusCode)
    }

    // Decodifica o JSON
    var pessoaJuridica models.PessoaJuridica
    if err := json.NewDecoder(resp.Body).Decode(&pessoaJuridica); err != nil {
        return nil, fmt.Errorf("erro ao decodificar JSON: %v", err)
    }

    return &pessoaJuridica, nil
}