package routes

import (
	"back/src/controllers"

	"github.com/gorilla/mux"
)

func loadDevicesControllers(router *mux.Router) {
	router.HandleFunc("/devices", controllers.GetRoomDevicesController).Methods("GET").Queries("roomId", "{roomId}")
	router.HandleFunc("/devices", controllers.GetDevicesController).Methods("GET")
	router.HandleFunc("/devices", controllers.PostDeviceController).Methods("POST")
	router.HandleFunc("/device/{deviceId}", controllers.GetDeviceController).Methods("GET")
	router.HandleFunc("/device/{deviceId}", controllers.PutDeviceController).Methods("PUT")
	router.HandleFunc("/device/{deviceId}/temperature", controllers.GetDeviceDataController).Methods("GET")
	router.HandleFunc("/device/{deviceId}", controllers.DeleteDeviceController).Methods("DELETE")
	router.HandleFunc("/device/{deviceId}/set_room_id", controllers.SetDeviceRoomController).Methods("PATCH")
}
