package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"consulta_api_gov/config"
	"consulta_api_gov/models"
)

func GetAuxilioEmergencialPorCPF(codigoBeneficiario string, pagina string) ([]models.AuxilioEmergencial, error) {
	var url = fmt.Sprintf("https://api.portaldatransparencia.gov.br/api-de-dados/auxilio-emergencial-por-cpf-ou-nis?codigoBeneficiario=%s&pagina=%s", codigoBeneficiario, pagina)

	fmt.Println("URL DA REQUEST: ", url)

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

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println("BODY DA RESPOSTA:", string(body))

	var result []models.AuxilioEmergencial
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
