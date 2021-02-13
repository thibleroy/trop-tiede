package routes

import (
	"fmt"
	"github.com/gorilla/mux"
)

func LoadRouters(router *mux.Router) {
	router.StrictSlash(true)
	loadUserRoutes(router)
	loadTracksRoutes(router)
	fmt.Println("routes loaded")
}
