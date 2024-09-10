package controller

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"integracaomobilemed/db"
	"integracaomobilemed/models"
	"io"
	"net/http"
	"os"
	//"github.com/gin-gonic/gin/internal/json"
)

// func AdicionarDados(c *gin.Context) {
// 	var dados models.Dados

// 	contentType := c.ContentType()

// 	switch contentType {
// 	case "application/json":
// 		if err := c.BindJSON(&dados); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 	case "application/xml":
// 		if err := c.BindXML(&dados); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Dados adicionados com sucesso!"})

// }

func AddDados(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")

	body, err := readRequestBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dados, formattedData, err := parseData(contentType, body)
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

func parseData(contentType string, body []byte) (models.Dados, string, error) {
	var dados models.Dados
	var formattedData string

	switch contentType {

	case "application/json":
		if err := json.Unmarshal(body, &dados); err != nil {
			return dados, "", err
		}
		formattedData = string(body)

	case "application/xml":
		if err := xml.Unmarshal(body, &dados); err != nil {
			return dados, "", err
		}
		formattedData = string(body)

	default:
		return dados, "", errors.New("unsupported content type")
	}

	return dados, formattedData, nil
}

func saveToDatabase(dados models.Dados) error {
	query := `INSERT INTO Dados (paciente, procedimento, plano, dados_clinicos)
		VALUES ($1, $2, $3, $4)`
	_, err := db.DB.Exec(query, dados.Paciente, dados.Procedimento, dados.Plano, dados.Dados_clinicos)
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
