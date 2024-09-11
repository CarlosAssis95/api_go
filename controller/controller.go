package controller

import (
	"encoding/json"
	"integracaomobilemed/db"
	"integracaomobilemed/models"
	"io"
	"net/http"
	"os"
	_"embed"
)

var (
	//go:embed query/insert_query.sql
	insert_query string
)

func AddDados(w http.ResponseWriter, r *http.Request) {
	body, err := readRequestBody(r)
	if err != nil {
		http.Error(w, "Falha ao ler o corpo da requisição", http.StatusInternalServerError)
		return
	}

	dados, formattedData, err := parseData(body)
	if err != nil {
		http.Error(w, "Falha ao fazer o parse dos dados", http.StatusInternalServerError)
		return
	}

	if err := saveToDatabase(dados); err != nil {
		http.Error(w, "Falha ao salvar no banco de dados", http.StatusInternalServerError)
		return
	}

	if err := saveToFile("SalvarDados.txt", formattedData); err != nil {
		http.Error(w, "Falha ao escrever no arquivo", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Dados adicionados com sucesso!"}`))
}

func readRequestBody(r *http.Request) ([]byte, error) {
	return io.ReadAll(r.Body)
}

func parseData(body []byte) (models.Dados, string, error) {
	var dados models.Dados

	if err := json.Unmarshal(body, &dados); err != nil {
		return dados, "", err
	}

	return dados, string(body), nil
}

func saveToDatabase(dados models.Dados) error {
	_, err := db.DB.Exec(insert_query, dados.Operacao, dados.NomePaciente, dados.PatientID, dados.DataNascimento, dados.Sexo, dados.AccessionNumber,
		dados.IdentificadorUnico, dados.Procedimento, dados.MedicoRadiologista, dados.CRMNR, dados.CRMUF, dados.Laudo, dados.LaudoRTF, dados.DataAssinatura,
		dados.DataExame, dados.Medico_solicitante, dados.Codigo_procedimento, dados.Tipo_exame, dados.Modalidade)
	return err
}

func saveToFile(fileName, data string) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.WriteString(data + "\n"); err != nil {
		return err
	}

	return nil
}
