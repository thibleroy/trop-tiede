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

const roomCollectionName = "room"
const temperatureCollectionName = "temperature"

func RetrieveRoom(id primitive.ObjectID) (*lib.IRoom, error) {
	roomToRetrieve := lib.IRoom{}
	err := lib.MyMusicAPIDB.Collection(roomCollectionName).FindOne(context.TODO(), bson.M{"resource.id": id}).Decode(&roomToRetrieve)
	if err != nil {
		return nil, err
	}
	return &roomToRetrieve, nil
}

func RetrieveRoomData(id primitive.ObjectID, startDate time.Time, endDate time.Time) (*[]lib.IRoomData, error) {
	roomDataToRetrieve := make([]lib.IRoomData, 0)
	cursor, _ := lib.MyMusicAPIDB.Collection(temperatureCollectionName).Find(context.TODO(), bson.M{"deviceid": id, "time": bson.M{"$lte": endDate, "$gte": startDate}})
	err := cursor.All(context.TODO(), &roomDataToRetrieve)
	if err != nil {
		return nil, err
	}
	fmt.Println("data retrieved", &roomDataToRetrieve)
	return &roomDataToRetrieve, nil
}

func AddRoomData(roomData lib.IRoomData) (*primitive.ObjectID, error) {
	fmt.Println("roomData", roomData)
	roomData.Resource = lib.NewResource()
	_, err := lib.MyMusicAPIDB.Collection(temperatureCollectionName).InsertOne(context.TODO(), roomData)
	if err != nil {
		return nil, err
	}
	return &roomData.Resource.ID, nil
}

func AddRoom(room lib.IRoom) (*primitive.ObjectID, error) {
	fmt.Println("room", room)
	room.Resource = lib.NewResource()
	_, err := lib.MyMusicAPIDB.Collection(roomCollectionName).InsertOne(context.TODO(), room)
	if err != nil {
		return nil, err
	}
	return &room.Resource.ID, nil
}

func UpdateRoom(room lib.IRoom) (*primitive.ObjectID, error) {
	updateTime := time.Now()
	room.Resource.UpdatedAt = updateTime
	_, err := lib.MyMusicAPIDB.Collection(roomCollectionName).UpdateOne(context.TODO(),bson.M{"resource.id": room.Resource.ID}, room)
	if err != nil {
		return nil, err
	}
	return &room.Resource.ID, nil
}

func RetrieveAllRooms() (*[]lib.IRoom, error) {
	retrievedRooms := make([]lib.IRoom, 0)
	cursor,_ := lib.MyMusicAPIDB.Collection(roomCollectionName).Find(context.TODO(), bson.M{})
	err := cursor.All(context.TODO(), &retrievedRooms)
	if err != nil {
		return nil, err
	}
	return &retrievedRooms, nil
}

func RemoveRoom(id primitive.ObjectID) (*primitive.ObjectID, error){
	_, err := lib.MyMusicAPIDB.Collection(roomCollectionName).DeleteOne(context.TODO(), bson.M{"resource.id": id})
	if err != nil {
		return nil, err
	}
	return &id, nil
}
