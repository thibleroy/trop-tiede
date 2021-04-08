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
)

func GetDeviceController (w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	Device, err := services.RetrieveDevice(objId)
	DeviceResponse := lib.IDeviceResponse{
		Device: Device,
	}
	lib.WriteToClient(w, lib.BuildResponse(DeviceResponse, err))
}

func GetDevicesController (w http.ResponseWriter, req *http.Request) {
	Devices, err := services.RetrieveAllDevices()
	DevicesReponse := lib.IDevicesResponse{Devices: Devices}
	lib.WriteToClient(w, lib.BuildResponse(DevicesReponse, err))
}

func PostDeviceController (w http.ResponseWriter, req *http.Request) {
	var Device lib.IDevice
	bodyDevice,_ := ioutil.ReadAll(req.Body)
	err := json.Unmarshal(bodyDevice, &Device)
	if err != nil {
		panic(err)
	}
	Device.Resource = lib.NewResource()
	returnId, serror := services.AddDevice(Device)
	if serror != nil {
		panic(serror)
	}
	w.Header().Add("Location", "http://" +req.Host + req.RequestURI + "/" + returnId.Hex())
	w.WriteHeader(201)
}

func DeleteDeviceController (w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	fmt.Println("id received for put", id)
	objId, errId := primitive.ObjectIDFromHex(id)
	if errId != nil {
		log.Fatal(errId)
	}
	_, _ = services.RemoveDevice(objId)
	w.WriteHeader(204)
}

func PutDeviceController (w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	fmt.Println("id received for put", id)
	objId, errId := primitive.ObjectIDFromHex(id)
	if errId != nil {
		log.Fatal(errId)
	}
	var Device lib.IDevice
	bodyTrack,_ := ioutil.ReadAll(req.Body)
	err := json.Unmarshal(bodyTrack, &Device)
	if err != nil {
		fmt.Println("error")
		log.Fatal(err)
	}
	Device.Resource.ID = objId
	returnId, dbError := services.UpdateDevice(Device)
	if dbError != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"`+dbError.Error()+`"}`))
		return
	}
	w.WriteHeader(204)
	w.Header().Add("Location", req.Host + req.RequestURI + "/" + returnId.Hex())
}
