package routes

import (
	"back/src/controllers"
	"github.com/gorilla/mux"
)
func loadRoomsControllers(router *mux.Router) {
	router.HandleFunc("/rooms", controllers.GetRoomsController).Methods("GET")
	router.HandleFunc("/rooms", controllers.PostRoomController).Methods("POST")
	router.HandleFunc("/room/{id}", controllers.GetRoomController).Methods("GET")
	router.HandleFunc("/room/{id}/temperature", controllers.GetRoomDataController).Methods("GET")
	router.HandleFunc("/room/{id}", controllers.PutRoomController).Methods("PUT")
	router.HandleFunc("/room/{id}", controllers.DeleteRoomController).Methods("DELETE")
}
