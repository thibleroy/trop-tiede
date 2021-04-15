package routes

import (
	"back/src/controllers"
	"github.com/gorilla/mux"
)
func loadRoomsControllers(router *mux.Router) {
	router.HandleFunc("/rooms", controllers.GetRoomsController).Methods("GET")
	router.HandleFunc("/rooms", controllers.PostRoomController).Methods("POST")
	router.HandleFunc("/room/{roomId}", controllers.GetRoomController).Methods("GET")
	router.HandleFunc("/room/{roomId}", controllers.PutRoomController).Methods("PUT")
	router.HandleFunc("/room/{roomId}", controllers.DeleteRoomController).Methods("DELETE")
	router.HandleFunc("/room/{roomId}/devices", controllers.GetRoomDevicesController).Methods("GET")
	router.HandleFunc("/room/{roomId}/devices", controllers.PostRoomDeviceController).Methods("POST")
	router.HandleFunc("/room/{roomId}/device/{deviceId}", controllers.DeleteRoomDeviceController).Methods("DELETE")
}
