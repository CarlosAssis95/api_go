package controller

import (
	"encoding/json"
	"integracaomobilemed/db"
	"integracaomobilemed/models"
	"io"
	"net/http"
	"os"
)
func AddDados(w http.ResponseWriter, r *http.Request) {

	body, err := readRequestBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dados, formattedData, err := parseData(body)
	if err != nil {
		if err.Error() == "unsupported content type" {
			http.Error(w, "Conteudo nao suportado", http.StatusUnsupportedMediaType)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	if err := saveToDatabase(dados); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func parseData(body []byte) (models.Dados, string, error) {
	var dados models.Dados
	var formattedData string

	if err := json.Unmarshal(body, &dados); err != nil {
		return dados, "", err
	}
	formattedData = string(body)

	return dados, formattedData, nil
}

func saveToDatabase(dados models.Dados) error {
	query := `INSERT INTO Dados (ds_operacao, ds_paciente, cd_paciente, dt_nascimento, ds_sexo, ds_accession_number, 
	nr_identificador, ds_procedimento, ds_medico, ds_crm_nr, ds_crm_uf, bb_laudo, bb_rtf, dt_assinatura, dt_data,
	medico_solicitante, codigo_procedimento, tipo_exame, modalidade)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)`
	_, err := db.DB.Exec(query, dados.Operacao, dados.NomePaciente, dados.PatientID, dados.DataNascimento, dados.Sexo, dados.AccessionNumber,
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
