package main

import (
	"back/lib"
	"back/lib/db"
	ttmqtt "back/lib/mqtt"
	"back/src/services"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strings"
	"time"
)

var tempStore []int

var temperatureReceivedHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	room := strings.Split(msg.Topic(),"/")[1]
	dataType := strings.Split(msg.Topic(),"/")[2]
	fmt.Printf("Received 1 message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	var roomData lib.IRoomData
	err := json.Unmarshal(msg.Payload(), &roomData)
	if err != nil {
		panic(err)
	}
	roomData.Time = time.Now()
	id, err := services.AddRoomData(roomData)
	if err != nil {
		panic(err)
	}
	fmt.Println("created id", id)
}

func main(){
	lib.Environment = lib.GetServerEnv()
	lib.MqttClient = ttmqtt.InitMqttClient(lib.Environment.MqttBrokerURL, lib.Environment.MqttBrokerPort, lib.Environment.MqttClientId, lib.Environment.MqttUsername, lib.Environment.MqttPassword)
	ttmqtt.ConnectMqttClient(lib.MqttClient)
	ttmqtt.Sub(lib.MqttClient, "#", temperatureReceivedHandler)
	// retrieves Mongo.Database instance
	lib.MyMusicAPIDB = db.InitDB(lib.Environment.MongoURL)
	select {}
}
