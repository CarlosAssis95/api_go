package models

type Dados struct {
	Paciente       string`json:"paciente" xml:"paciente"`
	Procedimento   string`json:"procedimento" xml:"procedimento"`
	Plano          string`json:"plano" xml:"plano"`
	Dados_clinicos string`json:"dados_clinicos" xml:"dados_clinicos"`
}
