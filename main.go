package main

import (
<<<<<<< Updated upstream
	"integracaomobilemed/db"
=======
	"fmt"
>>>>>>> Stashed changes
	"integracaomobilemed/routes"
)

func main() {

	rotas := routes.SetupRota()
	rotas.Run(":8080")
}
