package routes

import (
	"back/src/controllers"
	"github.com/gorilla/mux"
)
func loadDevicesControllers(router *mux.Router) {
	router.HandleFunc("/devices", controllers.GetDevicesController).Methods("GET")
	router.HandleFunc("/devices", controllers.PostDeviceController).Methods("POST")
	router.HandleFunc("/device/{id}", controllers.GetDeviceController).Methods("GET")
	router.HandleFunc("/device/{id}", controllers.PutDeviceController).Methods("PUT")
	router.HandleFunc("/device/{id}", controllers.DeleteDeviceController).Methods("DELETE")
}
