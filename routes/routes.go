package routes

import (
	"integracaomobilemed/controller"
	"integracaomobilemed/db"
	"integracaomobilemed/repository"

	"github.com/gin-gonic/gin"
)

<<<<<<< Updated upstream
func SetupRota() *gin.Engine {
	r := gin.Default()

	r.POST("/dados", controller.AddDados)

	return r
=======
func SetupRota() *mux.Router {
	router := mux.NewRouter()

	conectBanco := db.ConectaBanco()

	repo := repository.NewRepository(conectBanco)
	controller := controller.NewController(repo)

	router.HandleFunc("/dados", controller.AddDados).Methods("POST")

	return router
>>>>>>> Stashed changes
}
