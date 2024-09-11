package models

import "gopkg.in/guregu/null.v3"

type Dados struct {
	Operacao            string      `db:"ds_operacao" json:"Operacao"`
	NomePaciente        string      `db:"ds_paciente" json:"NomePaciente"`
	PatientID           int         `db:"cd_paciente" json:"PatientID"`
	DataNascimento      null.String `db:"dt_nascimento" json:"DataNascimento"`
	Sexo                null.String `db:"ds_sexo" json:"Sexo"`
	AccessionNumber     int         `db:"ds_accession_number" json:"AccessionNumber"`
	IdentificadorUnico  int         `db:"nr_identificador" json:"IdentificadorUnico"`
	Procedimento        null.String `db:"ds_procedimento" json:"Procedimento"`
	MedicoRadiologista  null.String `db:"ds_medico" json:"MedicoRadiologista"`
	CRMNR               null.String `db:"ds_crm_nr" json:"CRMNR"`
	CRMUF               null.String `db:"ds_crm_uf" json:"CRMUF"`
	Laudo               string      `db:"bb_laudo" json:"Laudo"`
	LaudoRTF            string      `db:"bb_rtf" json:"LaudoRTF"`
	DataAssinatura      null.String `db:"dt_assinatura" json:"DataAssinatura"`
	DataExame           null.String `db:"dt_data" json:"DataExame"`
	Medico_solicitante  string      `db:"medico_solicitante" json:"medicoSolicitante"`
	Codigo_procedimento string      `db:"codigo_procedimento" json:"codigoProcedimento"`
	Tipo_exame          string      `db:"tipo_exame" json:"tipoExame"`
	Modalidade          string      `db:"modalidade" json:"modalidade"`
}