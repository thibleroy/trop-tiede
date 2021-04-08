package main

import (
	"back/lib"
	ttmqtt "back/lib/mqtt"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func main(){
	lib.Environment = lib.GetServerEnv()
	lib.MqttClient = ttmqtt.InitMqttClient(lib.Environment.MqttBrokerURL, lib.Environment.MqttBrokerPort, lib.Environment.MqttClientId, lib.Environment.MqttUsername, lib.Environment.MqttPassword)
	ttmqtt.ConnectMqttClient(lib.MqttClient)
	for {
		data := lib.IRoomData{
			Resource:    lib.IResource{},
			Temperature: 20.5,
			Time:        time.Time{},
			RoomId:      primitive.ObjectID{},
			DeviceId:    primitive.ObjectID{},
		}
		payload,_ := json.Marshal(data)
		fmt.Println("*******************")
		fmt.Printf("%+v\n", data)
		fmt.Println("*******************")
		ttmqtt.Pub(lib.MqttClient, lib.Environment.MqttTemperatureTopic, string(payload))
		time.Sleep(5 * time.Second)
	}
}
