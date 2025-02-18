package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"consulta_api_gov/config"
	"consulta_api_gov/models"
)

func GetPessoaFisica(cpf string) (*models.PessoaFisica, error) {
	var url = fmt.Sprintf("%s/api-de-dados/pessoa-fisica?cpf=%s", config.LoadConfig().PortalAPIURL, cpf)

	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("chave-api-dados", config.LoadConfig().PortalAPIKey)

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Verifica o status HTTP da resposta
	if response.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("CPF não encontrado")
	} else if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro ao consultar a API: status %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result models.PessoaFisica
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}