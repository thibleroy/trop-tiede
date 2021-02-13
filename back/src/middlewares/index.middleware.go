package middlewares

import (
	"fmt"
	"github.com/gorilla/mux"
)

func LoadMiddlewares(router *mux.Router) {
	router.Use(authMiddleware)
	router.Use(loggingMiddleware)
	router.Use(headersMiddleware)
	fmt.Println("middlewares loaded")
}
