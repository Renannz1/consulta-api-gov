package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"consulta_api_gov/config"
	"consulta_api_gov/models"
)

func GetGarantiaSafra(cpf string, pagina string) ([]models.SafraBeneficiario, error) {
    config := config.LoadConfig()
    url := fmt.Sprintf("%s/api-de-dados/safra-codigo-por-cpf-ou-nis?codigo=%s&pagina=%s", config.PortalAPIURL, cpf, pagina)

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }

    req.Header.Add("chave-api-dados", config.PortalAPIKey)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("erro na requisição, status code: %d", resp.StatusCode)
    }

    var resultado []models.SafraBeneficiario
    if err := json.NewDecoder(resp.Body).Decode(&resultado); err != nil {
        return nil, err
    }

    return resultado, nil
}