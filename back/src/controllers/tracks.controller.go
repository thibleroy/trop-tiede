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

func GetTrackController (w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	fmt.Println("id received for get", id)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	track, getError := services.RetrieveTrack(objId)
	if getError != nil {
		w.WriteHeader(404)
		return
	}
	value, _ := json.Marshal(track)
	w.Write(value)
}

func GetTracksController (w http.ResponseWriter, req *http.Request) {
	tracks, err := services.RetrieveAllTracks()
	if err != nil {
		w.WriteHeader(404)
		return
	}
	value, _ := json.Marshal(tracks)
	w.Write(value)
}

func PostTrackController (w http.ResponseWriter, req *http.Request) {
	var track lib.IRoom
	bodyTrack,_ := ioutil.ReadAll(req.Body)
	err := json.Unmarshal(bodyTrack, &track)
	if err != nil {
		fmt.Println("error")
		log.Fatal(err)
	}
	track.Resource = lib.NewResource()
	returnId, err := services.AddTrack(track)
	w.Header().Add("Location", "http://" +req.Host + req.RequestURI + "/" + returnId.Hex())
	w.WriteHeader(201)
}

func PutTrackController (w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	fmt.Println("id received for put", id)
	objId, errId := primitive.ObjectIDFromHex(id)
	if errId != nil {
		log.Fatal(errId)
	}
	var track lib.IRoom
	bodyTrack,_ := ioutil.ReadAll(req.Body)
	err := json.Unmarshal(bodyTrack, &track)
	if err != nil {
		fmt.Println("error")
		log.Fatal(err)
	}
	track.Resource.ID = objId
	returnId, dbError := services.UpdateTrack(track)
	if dbError != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"`+dbError.Error()+`"}`))
		return
	}
	w.WriteHeader(204)
	w.Header().Add("Location", req.Host + req.RequestURI + "/" + returnId.Hex())
}

func DeleteTrackController (w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	fmt.Println("id received for put", id)
	objId, errId := primitive.ObjectIDFromHex(id)
	if errId != nil {
		log.Fatal(errId)
	}
	_, _ = services.RemoveTrack(objId)
	w.WriteHeader(204)
}
