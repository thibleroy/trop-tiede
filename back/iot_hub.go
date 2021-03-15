package main

import (
	"back/lib"
	"back/lib/db"
	ttmqtt "back/lib/mqtt"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strings"
)

var tempStore []int

var temperatureReceivedHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	room := strings.Split(msg.Topic(),"/")[1]
	dataType := strings.Split(msg.Topic(),"/")[2]
	fmt.Println("room", room)
	fmt.Println("datatype", dataType)
	fmt.Printf("Received 1 message: %s from topic: %s\n", msg.Payload(), msg.Topic())

}

func main(){
	lib.Environment = lib.GetServerEnv()
	lib.MqttClient = ttmqtt.InitMqttClient(lib.Environment.MqttBrokerURL, lib.Environment.MqttBrokerPort, lib.Environment.MqttClientId, lib.Environment.MqttUsername, lib.Environment.MqttPassword)
	ttmqtt.ConnectMqttClient(lib.MqttClient)
	ttmqtt.Sub(lib.MqttClient, "#", temperatureReceivedHandler)
	dbName := "trop-tiede"
	// retrieves Mongo.Database instance
	lib.MyMusicAPIDB = db.InitDB(lib.Environment.MongoURL, lib.Environment.MongoPort, dbName)
	select {}
}
