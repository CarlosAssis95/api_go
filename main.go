package main

import (
	"integracaomobilemed/db"
	"integracaomobilemed/routes"
)

func main() {
	db.ConectaBanco()

	r := routes.SetupRota()
	r.Run(":8080")
}
