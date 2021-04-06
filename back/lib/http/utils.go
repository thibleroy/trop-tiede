package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"strconv"
)

func InitWebServer(port int, router *mux.Router) {
	fmt.Println("Starting http server on port", port)
	handler := cors.Default().Handler(router)
	err := http.ListenAndServe(":" + strconv.Itoa(port), handler)
	if err != nil {
		log.Fatal(err)
	}
}
