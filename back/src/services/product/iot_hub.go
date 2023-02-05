package main

import (
	"back/lib"
	"back/lib/utils"
	"back/src/services"
	"encoding/json"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var tempStore []int

var temperatureReceivedHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received data: %s from device: %s\n", msg.Payload(), msg.Topic())
	var deviceData lib.IDeviceData
	err := json.Unmarshal(msg.Payload(), &deviceData)
	if err != nil {
		panic(err)
	}
	deviceData.Time = time.Now()
	id, _ := services.AddDeviceData(deviceData)
	fmt.Println("created id", id)
}

func main() {
	lib.Environment = utils.GetServerEnv()
	lib.MqttClient = utils.InitMqttClient(lib.Environment.MqttBrokerURL, lib.Environment.MqttBrokerPort, lib.Environment.MqttClientId, lib.Environment.MqttUsername, lib.Environment.MqttPassword)
	utils.ConnectMqttClient(lib.MqttClient)
	utils.Sub(lib.MqttClient, "/trop-tiede/#", temperatureReceivedHandler)
	dbName := "trop-tiede"
	// retrieves Mongo.Database instance
	lib.MyMusicAPIDB = utils.InitDB(lib.Environment.MongoURL, lib.Environment.MongoPort, dbName)
	select {}
}
