package controller

import (
	"integracaomobilemed/repository"
	"integracaomobilemed/services"
	"net/http"
	"strconv"
)

type ControllerInterface interface {
	AddDados(w http.ResponseWriter, r *http.Request)
}

type controllerStruct struct {
	repo repository.RepositoryInterface
}

func NewController(repo repository.RepositoryInterface) ControllerInterface {
	return &controllerStruct{
		repo: repo,
	}
}

func (c *controllerStruct) AddDados(w http.ResponseWriter, r *http.Request) {
	body, err := services.ReadRequestBody(r)
	if err != nil {
		http.Error(w, "Falha ao ler o corpo da requisição", http.StatusInternalServerError)
		return
	}

	dados, formattedData, err := services.ParseData(body)
	if err != nil {
		http.Error(w, "Falha ao fazer o parse dos dados", http.StatusInternalServerError)
		return
	}

	accessionNumber := strconv.Itoa(dados.AccessionNumber)
	if err := services.SaveToFile("SalvarDados.txt", formattedData, accessionNumber); err != nil {
		http.Error(w, "Falha ao escrever no arquivo", http.StatusInternalServerError)
		return
	}

	if err := c.repo.SaveDados(dados); err != nil {
		http.Error(w, "Falha ao salvar no banco de dados", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Dados adicionados com sucesso!"}`))
}
