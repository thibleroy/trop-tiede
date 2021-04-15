package controllers

import (
	"back/lib"
	"back/lib/utils"
	"back/src/services"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func GetUserController (w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	objId := utils.VerifyId(w, id)
	if objId != utils.EmptyId() {
		user, serviceStatus := services.RetrieveUser(objId)
		UserResponse := lib.IUserResponse{
			User: user,
		}
		utils.WriteToClient(w, utils.BuildResponse(UserResponse, utils.EmptyHeaders(), serviceStatus))
	}
}

func GetUsersController (w http.ResponseWriter, req *http.Request) {
	Users, serviceStatus := services.RetrieveAllUsers()
	UsersResponse := lib.IUsersResponse{Users: Users}
	utils.WriteToClient(w, utils.BuildResponse(UsersResponse, utils.EmptyHeaders(), serviceStatus))
}

func PostUserController (w http.ResponseWriter, req *http.Request) {
	var Users lib.IUser
	bodyUser,_ := ioutil.ReadAll(req.Body)
	err := json.Unmarshal(bodyUser, &Users)
	if err != nil {
		panic(err)
	}
	Users.Resource = utils.NewResource()
	returnId, serviceStatus := services.AddUser(Users)
	header := lib.IHeader{
		Key:   "Location",
		Value: "http://" + req.Host + req.RequestURI + "/" + returnId.Hex(),
	}
	utils.WriteToClient(w, utils.BuildResponse(nil, []lib.IHeader{header}, serviceStatus))
}

func DeleteUserController (w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	objId := utils.VerifyId(w, id)
	if objId != utils.EmptyId() {
		deleteId, serviceStatus := services.RemoveUser(objId)
		utils.WriteToClient(w, utils.BuildResponse(deleteId, utils.EmptyHeaders(), serviceStatus))
	}
}

func PutUserController (w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	objId := utils.VerifyId(w, id)
	if objId != utils.EmptyId() {
		var Users lib.IUser
		bodyTrack,_ := ioutil.ReadAll(req.Body)
		err := json.Unmarshal(bodyTrack, &Users)
		if err != nil {
			fmt.Println("error")
			log.Fatal(err)
		}
		Users.Resource.ID = objId
		returnId, serviceStatus := services.UpdateUser(Users)
		header := lib.IHeader{
			Key:   "Location",
			Value: "http://" + req.Host + req.RequestURI + "/" + returnId.Hex(),
		}
		utils.WriteToClient(w, utils.BuildResponse(nil, []lib.IHeader{header}, serviceStatus))
	}
}
