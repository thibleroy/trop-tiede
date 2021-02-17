package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func InitWebServer(port int, router *mux.Router) {
	fmt.Println("Starting http server on port", port)
	err := http.ListenAndServe(":" + strconv.Itoa(port), router)
	if err != nil {
		log.Fatal(err)
	}
}
