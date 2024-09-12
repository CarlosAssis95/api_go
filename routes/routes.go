package routes

import (
	"integracaomobilemed/controller"
	"integracaomobilemed/db"
	"integracaomobilemed/repository"

	"github.com/gorilla/mux"
)

func SetupRota() *mux.Router {
	router := mux.NewRouter()

	conectBanco := db.ConectaBanco()

	repo := repository.NewRepository(conectBanco)
	controller := controller.NewController(repo)

	router.HandleFunc("/dados", controller.AddDados).Methods("POST")

	return router
}
