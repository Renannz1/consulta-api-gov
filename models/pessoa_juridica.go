package models

type PessoaJuridica struct {
    CNPJ                     string `json:"cnpj"`
    RazaoSocial              string `json:"razaoSocial"`
    NomeFantasia             string `json:"nomeFantasia"`
    FavorecidoDespesas       bool   `json:"favorecidoDespesas"`
    PossuiContratacao        bool   `json:"possuiContratacao"`
    Convenios                bool   `json:"convenios"`
    FavorecidoTransferencias bool   `json:"favorecidoTransferencias"`
    SancionadoCEPIM          bool   `json:"sancionadoCEPIM"`
    SancionadoCEIS           bool   `json:"sancionadoCEIS"`
    SancionadoCNEP           bool   `json:"sancionadoCNEP"`
    SancionadoCEAF           bool   `json:"sancionadoCEAF"`
    ParticipanteLicitacao    bool   `json:"participanteLicitacao"`
    EmitiuNFe                bool   `json:"emitiuNFe"`
    BeneficiadoRenunciaFiscal bool  `json:"beneficiadoRenunciaFiscal"`
    IsentoImuneRenunciaFiscal bool  `json:"isentoImuneRenunciaFiscal"`
    HabilitadoRenunciaFiscal  bool  `json:"habilitadoRenunciaFiscal"`
}