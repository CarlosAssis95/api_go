package routes

import (
	"integracaomobilemed/controller"
	"net/http"
)

func SetupRota() {

	http.HandleFunc("/dados", controller.AddDados)

}
