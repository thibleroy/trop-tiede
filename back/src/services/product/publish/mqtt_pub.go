package main

import (
	"back/lib"
	utils "back/lib/utils"
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	lib.Environment = utils.GetServerEnv()
	lib.MqttClient = utils.InitMqttClient(lib.Environment.MqttBrokerURL, lib.Environment.MqttBrokerPort, lib.Environment.MqttClientId, lib.Environment.MqttUsername, lib.Environment.MqttPassword)
	utils.ConnectMqttClient(lib.MqttClient)
	for {
		data := lib.IDeviceData{
			Resource:    lib.IResource{},
			Temperature: 20.5,
			Time:        time.Time{},
			DeviceId:    primitive.ObjectID{},
		}
		payload, _ := json.Marshal(data)
		fmt.Println("*******************")
		fmt.Printf("%+v\n", data)
		fmt.Println("*******************")
		utils.Pub(lib.MqttClient, "/trop-tiede/", string(payload))
		time.Sleep(5 * time.Second)
	}
}
