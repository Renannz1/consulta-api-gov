package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"consulta_api_gov/config"
	"consulta_api_gov/models"
)

// GetExpenses busca os gastos públicos na API
func GetExpenses() ([]models.Expense, error) {
	conf := config.LoadConfig()
	url := fmt.Sprintf("%s/api-de-despesas", conf.PortalAPIURL) // Ajuste o endpoint correto

	// Criar requisição
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Adicionar headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-API-Key", conf.PortalAPIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Ler resposta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Converter JSON para struct
	var result models.APIExpensesResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result.Response, nil
}
