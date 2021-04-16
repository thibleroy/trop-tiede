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

const UserCollectionName = "User"

func RetrieveUser(id primitive.ObjectID) (lib.IUser, lib.IStatus) {
	userToRetrieve := lib.IUser{}
	err := lib.MyDB.Collection(UserCollectionName).FindOne(context.TODO(), bson.M{"resource.id": id}).Decode(&userToRetrieve)
	if err != nil {
		return lib.IUser{}, utils.FindError(UserCollectionName, 404)
	}
	return userToRetrieve, utils.FindSuccess(UserCollectionName, 200)
}

func AddUser(user lib.IUser) (primitive.ObjectID, lib.IStatus) {
	fmt.Println("User", user)
	user.Resource = utils.NewResource()
	_, err := lib.MyDB.Collection(UserCollectionName).InsertOne(context.TODO(), user)
	if err != nil {
		return primitive.ObjectID{}, utils.UpdateError("insert", UserCollectionName, 500)
	}
	return user.Resource.ID, utils.UpdateSuccess("insert", UserCollectionName, 201)
}

func UpdateUser(user lib.IUser) (primitive.ObjectID, lib.IStatus) {
	updateTime := time.Now()
	user.Resource.UpdatedAt = updateTime
	_, err := lib.MyDB.Collection(UserCollectionName).UpdateOne(context.TODO(),bson.M{"resource.id": user.Resource.ID}, user)
	if err != nil {
		return primitive.ObjectID{}, utils.UpdateError("update", UserCollectionName, 500)
	}
	return user.Resource.ID, utils.UpdateSuccess("update", UserCollectionName, 201)
}

func RetrieveAllUsers() ([]lib.IUser, lib.IStatus) {
	retrievedUsers := make([]lib.IUser, 0)
	cursor,_ := lib.MyDB.Collection(UserCollectionName).Find(context.TODO(), bson.M{})
	err := cursor.All(context.TODO(), &retrievedUsers)
	if err != nil {
		return nil, utils.FindError("Users", 404)
	}
	return retrievedUsers, utils.FindSuccess("Users", 200)
}

func RemoveUser(id primitive.ObjectID) (primitive.ObjectID, lib.IStatus){
	_, err := lib.MyDB.Collection(UserCollectionName).DeleteOne(context.TODO(), bson.M{"resource.id": id})
	if err != nil {
		return primitive.ObjectID{}, utils.UpdateError("remove", UserCollectionName, 500)
	}
	return id, utils.UpdateError("remove", UserCollectionName, 200)
}
