package routes

import (
	"integracaomobilemed/controller"

	"github.com/gorilla/mux"
)

func SetupRota() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/dados", controller.AddDados).Methods("POST")

    return router
}