package utils

import (
	"back/lib"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func WriteToClient (w http.ResponseWriter, response lib.IResponse) {
	for _, v := range response.Headers {
		w.Header().Add(v.Key, v.Value)
	}
	fmt.Println("status", response.Status.Code)
	w.WriteHeader(response.Status.Code)
	if response.Body.Value != nil {
		err := json.NewEncoder(w).Encode(response.Body)
		if err != nil {
			panic("Error writing to client " + err.Error())
		}
	}
}

func BuildResponse (body interface{}, headers []lib.IHeader, status lib.IStatus) (bodyValue lib.IResponse) {
	fmt.Println("status", status)
	return lib.IResponse{
		Status: status,
		Body: lib.IBody{
			Value:   body,
			Message: status.Message,
		},
		Headers: headers,
	}
}

func EmptyHeaders () []lib.IHeader {
	return make([]lib.IHeader, 0)
}

func VerifyId (w http.ResponseWriter, id string) primitive.ObjectID {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		status := lib.IStatus{
			Message: "Bad resource id",
			Code:    400,
		}
		WriteToClient(w, BuildResponse(id, EmptyHeaders(), status))
		return primitive.ObjectID{}
	}  else {
		return objId
	}
}

func EmptyId() primitive.ObjectID {
	return primitive.ObjectID{}
}
