package main

import (
	"back/lib"
	"back/lib/db"
	"back/src/services"
	"fmt"
)

func Init(){
	lib.Environment = lib.GetServerEnv()
	// retrieves Mongo.Database instance
	lib.MyMusicAPIDB = db.InitDB(lib.Environment.MongoURL)
}

func addDevice(){
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
	id, err := services.AddDevice(deviceToAdd)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("created device id", id)
	}
}

func addRoom(){
	roomToAdd := lib.IRoom{
		RoomDescription: lib.IRoomDescription{
			Description:  lib.IDescription{
				Name:    "Test Room name",
				Details: "Test Room details",
			},
		},
	}
	id, err := services.AddRoom(roomToAdd)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("created room id", id)
	}
}

func main(){
	Init()
	addDevice()
}
