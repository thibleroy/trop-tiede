package controllers

import (
	"back/lib"
	"back/lib/utils"
	"back/src/services"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"log"
	"net/http"
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
	roomsResponse := lib.IRoomsResponse{Rooms: rooms}
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
	room.DeviceIds = make([]primitive.ObjectID, 0)
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

func GetRoomDevicesController(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["roomId"]
	objId := utils.VerifyId(w, id)
	if objId != utils.EmptyId() {
		room, roomStatus := services.RetrieveRoom(objId)
		if roomStatus.Code == 404 {
			utils.WriteToClient(w, utils.BuildResponse(nil, utils.EmptyHeaders(), roomStatus))
		} else {
			devices := make([]lib.IDevice, 0)
			for _, deviceId := range room.DeviceIds {
				device, deviceStatus := services.RetrieveDevice(deviceId)
				if deviceStatus.Code == 404 {
					utils.WriteToClient(w, utils.BuildResponse(nil, utils.EmptyHeaders(), deviceStatus))
					return
				} else {
					devices = append(devices, device)
				}
			}
			utils.WriteToClient(w, utils.BuildResponse(devices, utils.EmptyHeaders(), utils.FindSuccess("devices", 200)))
		}
	}
}

func PostRoomDeviceController(w http.ResponseWriter, req *http.Request) {
	roomId := mux.Vars(req)["roomId"]
	roomObjId := utils.VerifyId(w, roomId)
	if roomObjId != utils.EmptyId() {
		room, roomStatus := services.RetrieveRoom(roomObjId)
		if roomStatus.Code == 404 {
			utils.WriteToClient(w, utils.BuildResponse(nil, utils.EmptyHeaders(), roomStatus))
		} else {
			var device lib.IDevice
			bodyDevice, _ := ioutil.ReadAll(req.Body)
			err := json.Unmarshal(bodyDevice, &device)
			if err != nil {
				fmt.Println("error")
				log.Fatal(err)
			}
			room.DeviceIds = append(room.DeviceIds, device.Resource.ID)
			updateId, status := services.UpdateRoom(room)
			utils.WriteToClient(w, utils.BuildResponse(updateId, utils.EmptyHeaders(), status))
		}
	}
}

func DeleteRoomDeviceController(w http.ResponseWriter, req *http.Request) {
	roomId := mux.Vars(req)["roomId"]
	roomObjId := utils.VerifyId(w, roomId)
	if roomObjId != utils.EmptyId() {
		room, roomStatus := services.RetrieveRoom(roomObjId)
		if roomStatus.Code == 404 {
			utils.WriteToClient(w, utils.BuildResponse(nil, utils.EmptyHeaders(), roomStatus))
		} else {
			deviceId := mux.Vars(req)["deviceId"]
			deviceObjId := utils.VerifyId(w, deviceId)
			if deviceObjId != utils.EmptyId() {
				newDevices := make([]primitive.ObjectID, 0)
				for _, val := range room.DeviceIds {
					if val != deviceObjId {
						newDevices = append(newDevices, val)
					}
				}
				room.DeviceIds = newDevices
				updateId, status := services.UpdateRoom(room)
				utils.WriteToClient(w, utils.BuildResponse(updateId, utils.EmptyHeaders(), status))
			}
		}
	}
}
