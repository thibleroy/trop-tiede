package routes

import (
	"github.com/gorilla/mux"
)

func LoadRouters(router *mux.Router) {
	router.StrictSlash(true)
	loadUserControllers(router)
	loadRoomsControllers(router)
	loadDevicesControllers(router)
}
