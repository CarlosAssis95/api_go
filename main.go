package main

import (
	"fmt"
	"integracaomobilemed/db"
	"integracaomobilemed/routes"
	"net/http"
)

func main() {
	db.ConectaBanco()

	router := routes.SetupRota()

	fmt.Println("Iniciando o servidor na porta 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
	}
}
