package controllers

import (
	"back/lib"
	"back/src/services"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetRoomController (w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	room, getError := services.RetrieveRoom(objId)
	if getError != nil {
		w.WriteHeader(404)
		return
	}
	value, _ := json.Marshal(room)
	w.Write(value)
}

func GetRoomsController (w http.ResponseWriter, req *http.Request) {
	rooms, err := services.RetrieveAllRooms()
	if err != nil {
		w.WriteHeader(404)
		return
	}
	a := lib.IRoomsResponse{Rooms: *rooms}
	value, merr := json.Marshal(a)
	if merr != nil {
		panic(merr)
	}
	fmt.Println("value", a.Rooms)
	w.Write(value)
}

func PostRoomDataController (w http.ResponseWriter, req *http.Request) {
	var roomData lib.IRoomData
	bodyRoom,_ := ioutil.ReadAll(req.Body)
	err := json.Unmarshal(bodyRoom, &roomData)
	if err != nil {
		fmt.Println("error")
		log.Fatal(err)
	}
	roomData.Resource = lib.NewResource()
	returnId, serror := services.AddRoomData(roomData)
	if serror != nil {
		panic(serror)
	}
	w.Header().Add("Location", "http://" +req.Host + req.RequestURI + "/" + returnId.Hex())
	w.WriteHeader(201)
}

func PostRoomController (w http.ResponseWriter, req *http.Request) {
	var room lib.IRoom
	bodyRoom,_ := ioutil.ReadAll(req.Body)
	err := json.Unmarshal(bodyRoom, &room)
	if err != nil {
		fmt.Println("error")
		log.Fatal(err)
	}
	room.Resource = lib.NewResource()
	returnId, serror := services.AddRoom(room)
	if serror != nil {
		panic(serror)
	}
	w.Header().Add("Location", "http://" +req.Host + req.RequestURI + "/" + returnId.Hex())
	w.WriteHeader(201)
}


func GetRoomDataController (w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	startDateInt,_ := strconv.ParseInt(req.URL.Query().Get("startDate"), 10, 64)
	endDateInt,_ := strconv.ParseInt(req.URL.Query().Get("endDate"), 10, 64)
	startDate := time.Unix(startDateInt, 0)
	endDate := time.Unix(endDateInt, 0)
	fmt.Println("id received for put", id)
	objId, errId := primitive.ObjectIDFromHex(id)
	if errId != nil {
		log.Fatal(errId)
	}
	roomData, dbError := services.RetrieveRoomData(objId, startDate, endDate)
	if dbError != nil {
		w.WriteHeader(404)
		return
	}
	a := lib.IRoomDataResponse{RoomData: *roomData}
	value, merr := json.Marshal(a)
	if merr != nil {
		panic(merr)
	}
	w.Write(value)
}

func DeleteRoomController (w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	fmt.Println("id received for put", id)
	objId, errId := primitive.ObjectIDFromHex(id)
	if errId != nil {
		log.Fatal(errId)
	}
	_, _ = services.RemoveRoom(objId)
	w.WriteHeader(204)
}

func PutRoomController (w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	fmt.Println("id received for put", id)
	objId, errId := primitive.ObjectIDFromHex(id)
	if errId != nil {
		log.Fatal(errId)
	}
	var room lib.IRoom
	bodyTrack,_ := ioutil.ReadAll(req.Body)
	err := json.Unmarshal(bodyTrack, &room)
	if err != nil {
		fmt.Println("error")
		log.Fatal(err)
	}
	room.Resource.ID = objId
	returnId, dbError := services.UpdateRoom(room)
	if dbError != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"`+dbError.Error()+`"}`))
		return
	}
	w.WriteHeader(204)
	w.Header().Add("Location", req.Host + req.RequestURI + "/" + returnId.Hex())
}
