package services

import (
	"back/lib"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "go.mongodb.org/mongo-driver/mongo"
	"math/rand"
	"time"
)

const roomCollectionName = "room"

func RetrieveRoom(id primitive.ObjectID) (*lib.IRoom, error) {
	//trackToRetrieve := lib.IRoom{}
	//err := lib.MyMusicAPIDB.Collection(roomCollectionName).FindOne(context.TODO(), bson.M{"resource.id": id}).Decode(&trackToRetrieve)
	//if err != nil {
	//	print("err", err.Error())
	//	return nil, err
	//}
	val := rand.Intn(5) + 20
	roomToRetrieve := lib.IRoom{
		Resource:    lib.IResource{
			ID:        primitive.ObjectID{},
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
		Name:        "Salon",
		Description: "Ceci est la description du salon.",
		Data:        lib.IRoomData{Temperature: val},
	}
	return &roomToRetrieve, nil
}

func AddRoom(room lib.IRoom) (*primitive.ObjectID, error) {
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
		"$set": room,
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
