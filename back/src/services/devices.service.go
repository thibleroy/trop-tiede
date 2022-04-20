package services

import (
	"back/lib"
	"back/lib/utils"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "go.mongodb.org/mongo-driver/mongo"
)

const DeviceCollectionName = "Device"

func RetrieveDevice(id primitive.ObjectID) (lib.IDevice, lib.IStatus) {
	deviceToRetrieve := lib.IDevice{}
	err := lib.MyMusicAPIDB.Collection(DeviceCollectionName).FindOne(context.TODO(), bson.M{"resource.id": id}).Decode(&deviceToRetrieve)
	if err != nil {
		return lib.IDevice{}, utils.FindError("Device", 404)
	}
	return deviceToRetrieve, utils.FindSuccess("Device", 200)
}

func AddDevice(device lib.IDevice) (primitive.ObjectID, lib.IStatus) {
	fmt.Println("Device", device)
	device.Resource = utils.NewResource()
	_, err := lib.MyMusicAPIDB.Collection(DeviceCollectionName).InsertOne(context.TODO(), device)
	if err != nil {
		return primitive.ObjectID{}, utils.UpdateError("insert", "Device", 500, err.Error())
	}
	return device.Resource.ID, utils.UpdateSuccess("insert", "Device", 201)
}

func UpdateDevice(device lib.IDevice) (primitive.ObjectID, lib.IStatus) {
	updateTime := time.Now()
	device.Resource.UpdatedAt = updateTime
	_, err := lib.MyMusicAPIDB.Collection(DeviceCollectionName).UpdateOne(context.TODO(), bson.M{"resource.id": device.Resource.ID}, bson.M{"$set": device})
	if err != nil {
		return primitive.ObjectID{}, utils.UpdateError("update", "Device", 500, err.Error())
	}
	return device.Resource.ID, utils.UpdateSuccess("update", "Device", 201)
}

func RetrieveAllDevices() ([]lib.IDevice, lib.IStatus) {
	retrievedDevices := make([]lib.IDevice, 0)
	cursor, _ := lib.MyMusicAPIDB.Collection(DeviceCollectionName).Find(context.TODO(), bson.M{})
	err := cursor.All(context.TODO(), &retrievedDevices)
	if err != nil {
		return nil, utils.FindError("Devices", 404)
	}
	return retrievedDevices, utils.FindSuccess("Devices", 200)
}

func RetrieveRoomDevices(roomId primitive.ObjectID) ([]lib.IDevice, lib.IStatus) {
	retrievedDevices := make([]lib.IDevice, 0)
	cursor, _ := lib.MyMusicAPIDB.Collection(DeviceCollectionName).Find(context.TODO(), bson.M{"roomid": roomId})
	err := cursor.All(context.TODO(), &retrievedDevices)
	if err != nil {
		return nil, utils.FindError("Devices", 404)
	}
	return retrievedDevices, utils.FindSuccess("Devices", 200)
}

func SetDeviceRoom(deviceId primitive.ObjectID, roomId primitive.ObjectID) {
	device, _ := RetrieveDevice(deviceId)
	device.RoomId = roomId
	UpdateDevice(device)
}

func AddDeviceData(deviceData lib.IDeviceData) (primitive.ObjectID, lib.IStatus) {
	fmt.Println("deviceData", deviceData)
	deviceData.Resource = utils.NewResource()
	_, err := lib.MyMusicAPIDB.Collection(temperatureCollectionName).InsertOne(context.TODO(), deviceData)
	if err != nil {
		return primitive.ObjectID{}, utils.UpdateError("insert", "room data", 500, err.Error())
	}
	return deviceData.Resource.ID, utils.UpdateSuccess("insert", "room data", 201)
}

func RemoveDevice(id primitive.ObjectID) (primitive.ObjectID, lib.IStatus) {
	_, err := lib.MyMusicAPIDB.Collection(DeviceCollectionName).DeleteOne(context.TODO(), bson.M{"resource.id": id})
	if err != nil {
		return primitive.ObjectID{}, utils.UpdateError("remove", DeviceCollectionName, 500, err.Error())
	}
	return id, utils.UpdateError("remove", DeviceCollectionName, 200, err.Error())
}

func RetrieveDeviceData(id primitive.ObjectID, startDate time.Time, endDate time.Time) ([]lib.IDeviceData, lib.IStatus) {
	deviceDataToRetrieve := make([]lib.IDeviceData, 0)
	fmt.Println("startdate", startDate.Unix(), "enddate", endDate.Unix())
	cursor, _ := lib.MyMusicAPIDB.Collection(temperatureCollectionName).Find(context.TODO(), bson.M{"deviceid": id, "time": bson.M{"$lte": endDate, "$gte": startDate}})
	err := cursor.All(context.TODO(), &deviceDataToRetrieve)
	if err != nil {
		return nil, utils.FindError("room data", 404)
	}
	fmt.Println("data retrieved", &deviceDataToRetrieve)
	return deviceDataToRetrieve, utils.FindSuccess("room data", 200)
}
