package routes

import (
	"github.com/gorilla/mux"
	"back/src/controllers"
)
func loadRoomsControllers(router *mux.Router) {
	router.HandleFunc("/rooms", controllers.GetRoomsController).Methods("GET")
	router.HandleFunc("/rooms", controllers.PostRoomController).Methods("POST")
	router.HandleFunc("/rooms/{id}", controllers.GetRoomController).Methods("GET")
	router.HandleFunc("/rooms/{id}", controllers.PutRoomController).Methods("PUT")
	router.HandleFunc("/rooms/{id}", controllers.DeleteRoomController).Methods("DELETE")
}
