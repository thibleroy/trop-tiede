package main

import (
	"back/lib"
	"back/lib/utils"
	"back/src/services"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func Init(){
	lib.Environment = utils.GetServerEnv()
	// retrieves Mongo.Database instance
	lib.MyDB = utils.InitDB(lib.Environment.MongoURL, lib.Environment.MongoName)
}

func addDevice() primitive.ObjectID{
	deviceToAdd := lib.IDevice{
		DeviceDescription: lib.IDeviceDescription{
			Position: lib.IPosition{
				Latitude:  22321,
				Longitude: 22332,
			},
			SerialNumber: "TestSerialNumber",
			Description:  lib.IDescription{
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

func addRoom() primitive.ObjectID{
	roomToAdd := lib.IRoom{
		RoomDescription: lib.IRoomDescription{
			Description:  lib.IDescription{
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

func addDeviceToRoom(deviceId primitive.ObjectID, roomId primitive.ObjectID) {
		room, _ := services.RetrieveRoom(roomId)
		fmt.Println("room retrieved", room)
		room.DeviceIds = append(room.DeviceIds, deviceId)
		updateId, status := services.UpdateRoom(room)
		fmt.Println("updated id", updateId)
		fmt.Println("status", status)
}

func main(){
	Init()
	rooms, status := services.RetrieveAllRooms()
	fmt.Println("rooms", rooms)
	fmt.Println("status", status)
	deviceId := addDevice()
	roomId := addRoom()
	fmt.Println("deviceId", deviceId)
	fmt.Println("roomId", roomId)
	addDeviceToRoom(deviceId, roomId)

	// objId, _ := primitive.ObjectIDFromHex("6071d09c43e80525657a2a29")
	// _, status := services.RetrieveDevice(objId)
	// if status.Code != 404 {
	// 	addDeviceData(objId)
	// } else {
	// 	fmt.Println("Error " + strconv.Itoa(status.Code), status.Message)
	// }
}
