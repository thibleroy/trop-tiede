package main

import (
	"back/lib"
	"back/lib/utils"
	"back/src/services"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strings"
)

var tempStore []int

var temperatureReceivedHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	room := strings.Split(msg.Topic(),"/")[1]
	dataType := strings.Split(msg.Topic(),"/")[2]
	fmt.Printf("Received 1 message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	
	roomToAdd := lib.IRoom{
		RoomDescription: lib.IRoomDescription{
			Description:  lib.IDescription{
				Name:    room,
				Details: dataType,
			},
		},
	}
	
	id, err := services.AddRoom(roomToAdd)
	if err.Code == 500 {
		panic(err)
	} else {
		fmt.Println("created room id", id)
	}
}

func main(){
	lib.Environment = utils.GetServerEnv()
	lib.MqttClient = utils.InitMqttClient(lib.Environment.MqttBrokerURL, lib.Environment.MqttBrokerPort, lib.Environment.MqttClientId, lib.Environment.MqttUsername, lib.Environment.MqttPassword)
	utils.ConnectMqttClient(lib.MqttClient)
	utils.Sub(lib.MqttClient, "#", temperatureReceivedHandler)
	// retrieves Mongo.Database instance
	lib.MyDB= utils.InitDB(lib.Environment.MongoURL, lib.Environment.MongoName)
	select {}
}