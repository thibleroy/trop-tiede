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
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func GetDeviceController(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["deviceId"]
	objId := utils.VerifyId(w, id)
	if objId != utils.EmptyId() {
		Device, serviceStatus := services.RetrieveDevice(objId)
		DeviceResponse := lib.IDeviceResponse{
			Device: Device,
		}
		utils.WriteToClient(w, utils.BuildResponse(DeviceResponse, utils.EmptyHeaders(), serviceStatus))
	}
}

func GetDevicesController(w http.ResponseWriter, req *http.Request) {
	devices, serviceStatus := services.RetrieveAllDevices()
	DevicesResponse := lib.IDevicesResponse{Devices: devices, Length: len(devices)}
	utils.WriteToClient(w, utils.BuildResponse(DevicesResponse, utils.EmptyHeaders(), serviceStatus))
}

func GetRoomDevicesController(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["roomId"]
	objId := utils.VerifyId(w, id)
	if objId != utils.EmptyId() {
		fmt.Println("roomid", objId)
		devices, serviceStatus := services.RetrieveRoomDevices(objId)
		DevicesResponse := lib.IDevicesResponse{Devices: devices, Length: len(devices)}
		utils.WriteToClient(w, utils.BuildResponse(DevicesResponse, utils.EmptyHeaders(), serviceStatus))
	}
}

func PostDeviceController(w http.ResponseWriter, req *http.Request) {
	var devices lib.IDevice
	bodyDevice, _ := ioutil.ReadAll(req.Body)
	err := json.Unmarshal(bodyDevice, &devices)
	if err != nil {
		panic(err)
	}
	devices.Resource = utils.NewResource()
	returnId, serviceStatus := services.AddDevice(devices)
	header := lib.IHeader{
		Key:   "Location",
		Value: "http://" + req.Host + req.RequestURI + "/" + returnId.Hex(),
	}
	utils.WriteToClient(w, utils.BuildResponse(nil, []lib.IHeader{header}, serviceStatus))
}

func DeleteDeviceController(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["deviceId"]
	objId := utils.VerifyId(w, id)
	if objId != utils.EmptyId() {
		deleteId, serviceStatus := services.RemoveDevice(objId)
		utils.WriteToClient(w, utils.BuildResponse(deleteId, utils.EmptyHeaders(), serviceStatus))
	}
}
func GetDeviceDataController(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["deviceId"]
	startDateInt, _ := strconv.ParseInt(req.URL.Query().Get("startDate"), 10, 64)
	endDateInt, _ := strconv.ParseInt(req.URL.Query().Get("endDate"), 10, 64)
	startDate := time.Unix(startDateInt, 0)
	endDate := time.Unix(endDateInt, 0)
	objId := utils.VerifyId(w, id)
	if objId != utils.EmptyId() {
		deviceData, serviceStatus := services.RetrieveDeviceData(objId, startDate, endDate)
		roomDataResponse := lib.IDeviceDataResponse{DeviceData: deviceData}
		utils.WriteToClient(w, utils.BuildResponse(roomDataResponse, utils.EmptyHeaders(), serviceStatus))
	}
}

func PutDeviceController(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["deviceId"]
	objId := utils.VerifyId(w, id)
	if objId != utils.EmptyId() {
		var devices lib.IDevice
		bodyTrack, _ := ioutil.ReadAll(req.Body)
		err := json.Unmarshal(bodyTrack, &devices)
		if err != nil {
			fmt.Println("error")
			log.Fatal(err)
		}
		devices.Resource.ID = objId
		returnId, serviceStatus := services.UpdateDevice(devices)
		header := lib.IHeader{
			Key:   "Location",
			Value: "http://" + req.Host + req.RequestURI + "/" + returnId.Hex(),
		}
		utils.WriteToClient(w, utils.BuildResponse(nil, []lib.IHeader{header}, serviceStatus))
	}
}

func SetDeviceRoomController(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["deviceId"]
	deviceId := utils.VerifyId(w, id)
	if deviceId != utils.EmptyId() {
		room := lib.IResource{}
		body, _ := ioutil.ReadAll(req.Body)
		err := json.Unmarshal(body, &room)
		if err != nil {
			fmt.Println("error")
			log.Fatal(err)
		}
		device, _ := services.RetrieveDevice(deviceId)
		device.RoomId = room.ID
		returnId, serviceStatus := services.UpdateDevice(device)
		header := lib.IHeader{
			Key:   "Location",
			Value: "http://" + req.Host + req.RequestURI + "/" + returnId.Hex(),
		}
		utils.WriteToClient(w, utils.BuildResponse(nil, []lib.IHeader{header}, serviceStatus))
	}
}
