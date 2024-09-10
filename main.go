package main

import (
	"integracaomobilemed/db"
	"integracaomobilemed/routes"
)

func main() {
	db.ConectaBanco()

	rotas := routes.SetupRota()
	rotas.Run(":8080")
}
