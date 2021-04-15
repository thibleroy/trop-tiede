package services

import (
	"back/lib"
	"back/lib/utils"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "go.mongodb.org/mongo-driver/mongo"
	"time"
)

const roomCollectionName = "room"
const temperatureCollectionName = "temperature"

func RetrieveRoom(id primitive.ObjectID) (lib.IRoom, lib.IStatus) {
	roomToRetrieve := lib.IRoom{}
	err := lib.MyMusicAPIDB.Collection(roomCollectionName).FindOne(context.TODO(), bson.M{"resource.id": id}).Decode(&roomToRetrieve)
	if err != nil {
		return lib.IRoom{}, utils.FindError(roomCollectionName, 404)
	}
	return roomToRetrieve, utils.FindSuccess(roomCollectionName, 200)
}

func AddRoom(room lib.IRoom) (primitive.ObjectID, lib.IStatus) {
	fmt.Println("room", room)
	room.Resource = utils.NewResource()
	_, err := lib.MyMusicAPIDB.Collection(roomCollectionName).InsertOne(context.TODO(), room)
	if err != nil {
		return primitive.ObjectID{}, utils.UpdateError("insert", roomCollectionName, 500)
	}
	return room.Resource.ID, utils.UpdateSuccess("insert", roomCollectionName, 201)
}

func UpdateRoom(room lib.IRoom) (primitive.ObjectID, lib.IStatus) {
	updateTime := time.Now()
	room.Resource.UpdatedAt = updateTime
	_, err := lib.MyMusicAPIDB.Collection(roomCollectionName).UpdateOne(context.TODO(),bson.M{"resource.id": room.Resource.ID}, bson.M{
		"$set": room,
	})
	if err != nil {
		fmt.Println("err", err)
		return primitive.ObjectID{}, utils.UpdateError("update", roomCollectionName, 500)
	}
	return room.Resource.ID, utils.UpdateSuccess("update", roomCollectionName, 201)
}

func RetrieveAllRooms() ([]lib.IRoom, lib.IStatus) {
	retrievedRooms := make([]lib.IRoom, 0)
	cursor,_ := lib.MyMusicAPIDB.Collection(roomCollectionName).Find(context.TODO(), bson.M{})
	err := cursor.All(context.TODO(), &retrievedRooms)
	if err != nil {
		return nil, utils.FindError("rooms", 404)
	}
	return retrievedRooms, utils.FindSuccess("rooms", 200)
}

func RemoveRoom(id primitive.ObjectID) (primitive.ObjectID, lib.IStatus){
	_, err := lib.MyMusicAPIDB.Collection(roomCollectionName).DeleteOne(context.TODO(), bson.M{"resource.id": id})
	if err != nil {
		return primitive.ObjectID{}, utils.UpdateError("remove", roomCollectionName, 500)
	}
	return id, utils.UpdateError("remove", roomCollectionName, 200)
}
