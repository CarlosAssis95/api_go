package services

import (
	"encoding/json"
	"integracaomobilemed/models"
	"io"
	"net/http"
	"os"
	"time"
)

func ReadRequestBody(r *http.Request) ([]byte, error) {
	return io.ReadAll(r.Body)	
}

func SaveToFile(fileName string, data string, accessionNumber string) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	now := time.Now().Format("02-01-2006 15:04:05")
	formattedData := now + " accessionNumber " + accessionNumber + "  " + data

	if _, err = file.WriteString(formattedData + "\n"); err != nil {
		return err
	}

	return nil
}

func ParseData(body []byte) (models.Dados, string, error) {
	var dados models.Dados

	if err := json.Unmarshal(body, &dados); err != nil {
		return dados, "", err
	}

	return dados, string(body), nil
}
