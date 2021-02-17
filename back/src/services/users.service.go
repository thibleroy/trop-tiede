package services

import (
	"back/lib"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const userCollectionName = "user"

func RetrieveUser(id primitive.ObjectID) (*lib.IUser, error) {
	userToRetrieve := lib.IUser{}
	err := lib.MyMusicAPIDB.Collection(userCollectionName).FindOne(context.TODO(), bson.M{"resource.id": id}).Decode(&userToRetrieve)
	if err != nil {
		print("err", err.Error())
		return nil, err
	}
	return &userToRetrieve, nil
}

func UpdateUser(user lib.IUser) (*primitive.ObjectID, error) {
	user.Resource.UpdatedAt = time.Now()
	update := bson.M{
		"$set": user,
	}
	_, err := lib.MyMusicAPIDB.Collection(userCollectionName).UpdateOne(context.TODO(),bson.M{"resource.id": user.Resource.ID}, update)
	if err != nil {
		return nil, err
	}
	return &user.Resource.ID, nil
}

func UserSignup(user lib.IUser) (*primitive.ObjectID, error){
	fmt.Println("user signup", &user)
	_, err := lib.MyMusicAPIDB.Collection(userCollectionName).InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	return &user.Resource.ID, nil
}

func UserLogin(user lib.IUser) (*string, error) {
	fmt.Println("user login", &user)
	userToRetrieve := lib.IUser{}
	err := lib.MyMusicAPIDB.Collection(userCollectionName).FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&userToRetrieve)
	if err != nil {
		print("err", err.Error())
		return nil, err
	}
	return &userToRetrieve.Password, nil
}

func RemoveUser(id primitive.ObjectID) (*primitive.ObjectID, error){
	_, err := lib.MyMusicAPIDB.Collection(userCollectionName).DeleteOne(context.TODO(), bson.M{"resource.id": id})
	if err != nil {
		return nil, err
	}
	return &id, nil
}
