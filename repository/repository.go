package repository

import (
	"database/sql"
	_ "embed"
	"integracaomobilemed/models"
)

var (
	//go:embed query/insert_query.sql
	insert_query string
)

type RepositoryInterface interface {
	SaveDados(dados models.Dados) error
}

type repositoryStruct struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) RepositoryInterface {
	return &repositoryStruct{
		db: db,
	}
}

func (r *repositoryStruct) SaveDados(dados models.Dados) error {
	return r.saveToDatabase(dados)
}

func (r *repositoryStruct) saveToDatabase(dados models.Dados) error {
	_, err := r.db.Exec(insert_query, dados.Operacao, dados.NomePaciente, dados.PatientID, dados.DataNascimento,
		dados.Sexo, dados.AccessionNumber, dados.IdentificadorUnico, dados.Procedimento, dados.MedicoRadiologista,
		dados.CRMNR, dados.CRMUF, dados.Laudo, dados.LaudoRTF, dados.DataAssinatura, dados.DataExame, dados.Medico_solicitante,
		dados.Codigo_procedimento, dados.Tipo_exame, dados.Modalidade)
	return err
}
