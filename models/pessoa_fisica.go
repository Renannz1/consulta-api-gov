package models

type PessoaFisica struct {
	Cpf                             string `json:"cpf"`
	Nome                            string `json:"nome"`
	Servidor                        bool   `json:"servidor"`
	Contratado                      bool   `json:"contratado"`
	FavorecidoBolsaFamilia          bool   `json:"favorecidoBolsaFamilia"`
	FavorecidoSafra                 bool   `json:"favorecidoSafra"`
	AuxilioEmergencial              bool   `json:"auxilioEmergencial"`
	FavorecidoAuxilioBrasil         bool   `json:"favorecidoAuxilioBrasil"`
	FavorecidoNovoBolsaFamilia      bool   `json:"favorecidoNovoBolsaFamilia"`
	FavorecidoDespesas              bool   `json:"favorecidoDespesas"`
	BeneficiarioDiarias             bool   `json:"beneficiarioDiarias"`
	Permissionario                  bool   `json:"permissionario"`
	SancionadoCEIS                  bool   `json:"sancionadoCEIS"`
	SancionadoCNEP                  bool   `json:"sancionadoCNEP"`
	SancionadoCEAF                  bool   `json:"sancionadoCEAF"`
	PortadorCPDC                    bool   `json:"portadorCPDC"`
	PortadorCPGF                    bool   `json:"portadorCPGF"`
	FavorecidoPeti                  bool   `json:"favorecidoPeti"`
	FavorecidoSeguroDefeso          bool   `json:"favorecidoSeguroDefeso"`
	FavorecidoBpc                   bool   `json:"favorecidoBpc"`
	FavorecidoTransferencias        bool   `json:"favorecidoTransferencias"`
	FavorecidoCPCC                  bool   `json:"favorecidoCPCC"`
	FavorecidoCPDC                  bool   `json:"favorecidoCPDC"`
	FavorecidoCPGF                  bool   `json:"favorecidoCPGF"`
	ParticipanteLicitacao           bool   `json:"participanteLicitacao"`
	ServidorInativo                 bool   `json:"servidorInativo"`
	PensionistaOuRepresentanteLegal bool   `json:"pensionistaOuRepresentanteLegal"`
	InstituidorPensao               bool   `json:"instituidorPensao"`
	FavorecidoAuxilioReconstrucao   bool   `json:"favorecidoAuxilioReconstrucao"`
}
