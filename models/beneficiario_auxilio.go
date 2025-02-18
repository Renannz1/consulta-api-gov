package models

type Beneficiario struct {
	CpfFormatado string `json:"cpfFormatado"`
	Nis          string `json:"nis"`
	Nome         string `json:"nome"`
}

type Municipio struct {
	CodigoIBGE   string `json:"codigoIBGE"`
	NomeIBGE     string `json:"nomeIBGE"`
	CodigoRegiao string `json:"codigoRegiao"`
	NomeRegiao   string `json:"nomeRegiao"`
	Pais         string `json:"pais"`
	Uf           struct {
		Sigla string `json:"sigla"`
		Nome  string `json:"nome"`
	} `json:"uf"`
}

type AuxilioEmergencial struct {
	Id                              int          `json:"id"`
	MesDisponibilizacao             string       `json:"mesDisponibilizacao"`
	Beneficiario                    Beneficiario `json:"beneficiario"`
	ResponsavelAuxilioEmergencial   Beneficiario `json:"responsavelAuxilioEmergencial"`
	Municipio                       Municipio    `json:"municipio"`
	SituacaoAuxilioEmergencial      string       `json:"situacaoAuxilioEmergencial"`
	EnquadramentoAuxilioEmergencial string       `json:"enquadramentoAuxilioEmergencial"`
	Valor                           float64      `json:"valor"`
	NumeroParcela                   string       `json:"numeroParcela"`
}
