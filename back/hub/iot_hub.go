package main

import (
	"back/lib"
	"back/lib/utils"
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
	lib.Environment = utils.GetServerEnv()
	lib.MqttClient = utils.InitMqttClient(lib.Environment.MqttBrokerURL, lib.Environment.MqttBrokerPort, lib.Environment.MqttClientId, lib.Environment.MqttUsername, lib.Environment.MqttPassword)
	utils.ConnectMqttClient(lib.MqttClient)
	utils.Sub(lib.MqttClient, "#", temperatureReceivedHandler)
	dbName := "trop-tiede"
	// retrieves Mongo.Database instance
	lib.MyMusicAPIDB = utils.InitDB(lib.Environment.MongoURL, lib.Environment.MongoPort, dbName)
	select {}
}
