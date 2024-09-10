package controller

import (
	"encoding/json"
	"encoding/xml"
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

func AdicionarDados(c *gin.Context) {
	var dados models.Dados

	contentType := c.ContentType()

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	fileName := "SalvarDados.txt"
	var formattedData string

	switch contentType {
	
	case "application/json":
		if err := json.Unmarshal(body, &dados); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		formattedData = string(body)
	
	case "application/xml":
		if err := xml.Unmarshal(body, &dados); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		formattedData = string(body)
	}

	query := `INSERT INTO Dados (paciente, procedimento, plano, dados_clinicos)
		VALUES ($1, $2, $3, $4)`

	_, err = db.DB.Exec(query, dados.Paciente, dados.Procedimento, dados.Plano, dados.Dados_clinicos)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer file.Close()

	_, err = file.WriteString(formattedData + "\n")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write to file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Dados adicionados com sucesso!"})

}
