package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"back/src/lib"
	"back/src/services"
	"io/ioutil"
	"log"
	"net/http"
)

func PostUserController (w http.ResponseWriter, req *http.Request) {
	var user lib.IUser
	bodyUser,_ := ioutil.ReadAll(req.Body)
	err := json.Unmarshal(bodyUser, &user)
	if err != nil {
		fmt.Println("error")
		log.Fatal(err)
	}
	user.Resource = lib.NewResource()
	user.Password = lib.GetHash([]byte(user.Password))
	returnId, err := services.UserSignup(user)
	w.Header().Add("Location", "http://" + req.Host + req.RequestURI + "/" + returnId.Hex())
	w.WriteHeader(201)
}

func GetUserController (w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	fmt.Println("id received for get", id)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	user, getError := services.RetrieveUser(objId)
	if getError != nil {
		w.WriteHeader(404)
		return
	}
	value, _ := json.Marshal(user)
	w.Write(value)
}

func PutUserController (w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	fmt.Println("id received for put", id)
	objId, errId := primitive.ObjectIDFromHex(id)
	if errId != nil {
		log.Fatal(errId)
	}
	var user lib.IUser
	bodyUser,_ := ioutil.ReadAll(req.Body)
	err := json.Unmarshal(bodyUser, &user)
	if err != nil {
		fmt.Println("error")
		log.Fatal(err)
	}
	user.Resource.ID = objId
	returnId, dbError := services.UpdateUser(user)
	if dbError!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"`+dbError.Error()+`"}`))
		return
	}
	w.WriteHeader(204)
	w.Header().Add("Location", req.Host + req.RequestURI + "/" + returnId.Hex())
}

func DeleteUserController (w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	fmt.Println("id received for delete", id)
	objId, errId := primitive.ObjectIDFromHex(id)
	if errId != nil {
		log.Fatal(errId)
	}
	_, _ = services.RemoveUser(objId)
	w.WriteHeader(204)
}

