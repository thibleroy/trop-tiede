package routes

import (
	"github.com/gorilla/mux"
	"back/src/controllers"
)
func loadSessionRoutes(router *mux.Router) {
	router.HandleFunc("/tracks", controllers.GetTracksController).Methods("GET")
	router.HandleFunc("/tracks", controllers.PostTrackController).Methods("POST")
	router.HandleFunc("/tracks/{id}", controllers.GetTrackController).Methods("GET")
	router.HandleFunc("/tracks/{id}", controllers.PutTrackController).Methods("PUT")
	router.HandleFunc("/tracks/{id}", controllers.DeleteTrackController).Methods("DELETE")
}

