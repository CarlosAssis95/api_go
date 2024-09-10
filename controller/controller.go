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

	"github.com/gin-gonic/gin"
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

func AddDados(c *gin.Context) {
	contentType := c.ContentType()

	body, err := readRequestBody(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	dados, formattedData, err := parseData(contentType, body)
	if err != nil {
		if err.Error() == "unsupported content type" {
			c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": "Conteudo não suportado"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	if err := saveToDatabase(dados); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := saveToFile("SalvarDados.txt", formattedData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Dados adicionados com sucesso!"})
}

func readRequestBody(c *gin.Context) ([]byte, error) {
	body, err := io.ReadAll(c.Request.Body)
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
