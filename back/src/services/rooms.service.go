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

func RetrieveRoom(id primitive.ObjectID) (*lib.IRoom, error) {
	roomToRetrieve := lib.IRoom{}
	err := lib.MyMusicAPIDB.Collection(roomCollectionName).FindOne(context.TODO(), bson.M{"resource.id": id}).Decode(&roomToRetrieve)
	if err != nil {
		return nil, err
	}
	return &roomToRetrieve, nil
}

func RetrieveRoomData(id primitive.ObjectID, startDate primitive.DateTime, endDate primitive.DateTime) (*lib.IRoom, error) {
	roomToRetrieve := lib.IRoom{}
	cursor, _ := lib.MyMusicAPIDB.Collection(roomCollectionName).Find(context.TODO(), bson.M{"resource.id": id, "data.time": bson.M{"$lte": endDate, "$gte": startDate}})
	err := cursor.All(context.TODO(), &roomToRetrieve)
	if err != nil {
		return nil, err
	}
	fmt.Println("data retrieved", roomToRetrieve)
	return &roomToRetrieve, nil
}

func AddRoomData(room lib.IRoom) (*primitive.ObjectID, error) {
	fmt.Println("obj", room)
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
	update := bson.M{
		"$push": room.Data,
	}
	_, err := lib.MyMusicAPIDB.Collection(roomCollectionName).UpdateOne(context.TODO(),bson.M{"resource.id": room.Resource.ID}, update)
	if err != nil {
		return nil, err
	}
	return &room.Resource.ID, nil
}

func RetrieveAllRooms() (*[]lib.IRoom, error) {
	var retrievedRooms []lib.IRoom
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
