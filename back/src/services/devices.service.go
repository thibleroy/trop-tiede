package services

import (
	"back/lib"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "go.mongodb.org/mongo-driver/mongo"
	"time"
)

const deviceCollectionName = "device"

func RetrieveDevice(id primitive.ObjectID) (lib.IDevice, error) {
	deviceToRetrieve := lib.IDevice{}
	err := lib.MyMusicAPIDB.Collection(deviceCollectionName).FindOne(context.TODO(), bson.M{"resource.id": id}).Decode(&deviceToRetrieve)
	if err != nil {
		return lib.IDevice{}, err
	}
	return deviceToRetrieve, nil
}

func AddDevice(device lib.IDevice) (primitive.ObjectID, error) {
	fmt.Println("device", device)
	device.Resource = lib.NewResource()
	_, err := lib.MyMusicAPIDB.Collection(deviceCollectionName).InsertOne(context.TODO(), device)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return device.Resource.ID, nil
}

func UpdateDevice(device lib.IDevice) (primitive.ObjectID, error) {
	updateTime := time.Now()
	device.Resource.UpdatedAt = updateTime
	_, err := lib.MyMusicAPIDB.Collection(deviceCollectionName).UpdateOne(context.TODO(),bson.M{"resource.id": device.Resource.ID}, device)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return device.Resource.ID, nil
}

func RetrieveAllDevices() ([]lib.IDevice, error) {
	retrievedDevices := make([]lib.IDevice, 0)
	cursor,_ := lib.MyMusicAPIDB.Collection(deviceCollectionName).Find(context.TODO(), bson.M{})
	err := cursor.All(context.TODO(), &retrievedDevices)
	if err != nil {
		return nil, err
	}
	return retrievedDevices, nil
}

func RemoveDevice(id primitive.ObjectID) (primitive.ObjectID, error){
	_, err := lib.MyMusicAPIDB.Collection(deviceCollectionName).DeleteOne(context.TODO(), bson.M{"resource.id": id})
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return id, nil
}
