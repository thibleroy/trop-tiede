package controllers

import (
	"back/lib"
	"back/lib/utils"
	"back/src/services"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetRoomController(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["roomId"]
	objId := utils.VerifyId(w, id)
	if objId != utils.EmptyId() {
		room, serviceStatus := services.RetrieveRoom(objId)
		roomResponse := lib.IRoomResponse{
			Room: room,
		}
		utils.WriteToClient(w, utils.BuildResponse(roomResponse, utils.EmptyHeaders(), serviceStatus))
	}
}

func GetRoomsController(w http.ResponseWriter, req *http.Request) {
	rooms, serviceStatus := services.RetrieveAllRooms()
	roomsResponse := lib.IRoomsResponse{Rooms: rooms, Length: len(rooms)}
	fmt.Println("rooms", roomsResponse)
	utils.WriteToClient(w, utils.BuildResponse(roomsResponse, utils.EmptyHeaders(), serviceStatus))
}

func PostRoomController(w http.ResponseWriter, req *http.Request) {
	var room lib.IRoom
	bodyRoom, _ := ioutil.ReadAll(req.Body)
	err := json.Unmarshal(bodyRoom, &room)
	if err != nil {
		panic(err)
	}
	room.Resource = utils.NewResource()
	returnId, serviceStatus := services.AddRoom(room)
	header := lib.IHeader{
		Key:   "Location",
		Value: "http://" + req.Host + req.RequestURI + "/" + returnId.Hex(),
	}
	utils.WriteToClient(w, utils.BuildResponse(nil, []lib.IHeader{header}, serviceStatus))
}

func DeleteRoomController(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["roomId"]
	objId := utils.VerifyId(w, id)
	if objId != utils.EmptyId() {
		deleteId, serviceStatus := services.RemoveRoom(objId)
		utils.WriteToClient(w, utils.BuildResponse(deleteId, utils.EmptyHeaders(), serviceStatus))
	}
}

func PutRoomController(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["roomId"]
	objId := utils.VerifyId(w, id)
	if objId != utils.EmptyId() {
		var room lib.IRoom
		bodyRoom, _ := ioutil.ReadAll(req.Body)
		err := json.Unmarshal(bodyRoom, &room)
		if err != nil {
			fmt.Println("error")
			log.Fatal(err)
		}
		room.Resource.ID = objId
		returnId, serviceStatus := services.UpdateRoom(room)

		header := lib.IHeader{
			Key:   "Location",
			Value: "http://" + req.Host + req.RequestURI + "/" + returnId.Hex(),
		}
		utils.WriteToClient(w, utils.BuildResponse(nil, []lib.IHeader{header}, serviceStatus))
	}
}
