package models

type BeneficiarioSafra struct {
	CPFFormatado string `json:"cpfFormatado"`
	NIS         string `json:"nis"`
	Nome        string `json:"nome"`
}

type UF struct {
	Sigla string `json:"sigla"`
	Nome  string `json:"nome"`
}

type Mun struct {
	CodigoIBGE   string `json:"codigoIBGE"`
	NomeIBGE     string `json:"nomeIBGE"`
	CodigoRegiao string `json:"codigoRegiao"`
	NomeRegiao   string `json:"nomeRegiao"`
	Pais         string `json:"pais"`
	UF          UF     `json:"uf"`
}

type SafraBeneficiario struct {
	ID               int              `json:"id"`
	BeneficiarioSafra BeneficiarioSafra `json:"beneficiarioSafra"`
	DataMesReferencia string           `json:"dataMesReferencia"`
	Mun        Mun         `json:"municipio"`
	Valor            float64           `json:"valor"`
}