package main

import (
	"back/lib"
	"back/lib/utils"
	"back/src/services"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Init() {
	lib.Environment = utils.GetServerEnv()
	dbName := "trop-tiede"
	// retrieves Mongo.Database instance
	lib.MyMusicAPIDB = utils.InitDB(lib.Environment.MongoURL, lib.Environment.MongoPort, dbName)
}

func addDevice() primitive.ObjectID {
	deviceToAdd := lib.IDevice{
		DeviceDescription: lib.IDeviceDescription{
			Position: lib.IPosition{
				Latitude:  22321,
				Longitude: 22332,
			},
			SerialNumber: "TestSerialNumber",
			Description: lib.IDescription{
				Name:    "Test Device name",
				Details: "Test Device details",
			},
		},
	}
	id, status := services.AddDevice(deviceToAdd)
	fmt.Println("status", status)
	if status.Code != 201 {
		panic(status)
	} else {
		fmt.Println("created device id", id)
	}
	return id
}

func addRoom() primitive.ObjectID {
	roomToAdd := lib.IRoom{
		RoomDescription: lib.IRoomDescription{
			Description: lib.IDescription{
				Name:    "Test Room name",
				Details: "Test Room details",
			},
		},
	}
	id, err := services.AddRoom(roomToAdd)
	if err.Code == 500 {
		panic(err)
	} else {
		fmt.Println("created room id", id)
	}
	return id
}

func addDeviceData(deviceId primitive.ObjectID) {
	deviceData := lib.IDeviceData{
		Resource:    lib.IResource{},
		Temperature: 22,
		Time:        time.Now(),
		DeviceId:    deviceId,
	}
	id, status := services.AddDeviceData(deviceData)
	fmt.Println("device data inserted id", id)
	fmt.Println("status", status)
}

func setDeviceRoom(deviceId primitive.ObjectID, roomId primitive.ObjectID) {
	device, _ := services.RetrieveDevice(deviceId)
	fmt.Println("device to update", device)
	device.RoomId = roomId
	updateId, status := services.UpdateDevice(device)
	fmt.Println("updated id", updateId)
	fmt.Println("status", status)
}

func main() {
	Init()
	deviceId := addDevice()
	roomId := addRoom()
	fmt.Println("deviceId", deviceId)
	fmt.Println("roomId", roomId)
	setDeviceRoom(deviceId, roomId)

	//objId, _ := primitive.ObjectIDFromHex("6071d09c43e80525657a2a29")
	//_, status := services.RetrieveDevice(objId)
	//if status.Code != 404 {
	//	addDeviceData(objId)
	//} else {
	//	fmt.Println("Error " + strconv.Itoa(status.Code), status.Message)
	//}
}
