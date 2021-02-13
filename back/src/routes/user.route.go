package routes

import (
	"github.com/gorilla/mux"
	"back/src/controllers"
)
func loadUserRoutes(router *mux.Router) {
	router.HandleFunc("/users", controllers.PostUserController).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.GetUserController).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.PutUserController).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUserController).Methods("DELETE")
}


