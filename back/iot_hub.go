package main

import (
	"back/lib"
	"back/lib/db"
	ttmqtt "back/lib/mqtt"
	"back/src/services"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strconv"
	"strings"
	"time"
)

var tempStore []int

var temperatureReceivedHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Println(strings.Split(msg.Topic(),"/"))
	room := strings.Split(msg.Topic(),"/")[1]
	dataType := strings.Split(msg.Topic(),"/")[2]
	fmt.Println("room", room)
	fmt.Println("datatype", dataType)
	fmt.Printf("Received 1 message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	num, _ := strconv.ParseFloat(string(msg.Payload()),64)
	currentTemp := lib.IRoomData{
		Temperature: num,
		Time:        time.Now(),
	}
	currentRoom := lib.IRoom{
		Name:        room,
		Description: "Room description of " + room,
		Data:        currentTemp,
	}
	_, err := services.AddRoomData(currentRoom)
	if err != nil {
		panic(err)
	}
}

func main(){
	lib.Environment = lib.GetServerEnv()
	lib.MqttClient = ttmqtt.InitMqttClient(lib.Environment.MqttBrokerURL, lib.Environment.MqttBrokerPort, lib.Environment.MqttClientId, lib.Environment.MqttUsername, lib.Environment.MqttPassword)
	ttmqtt.ConnectMqttClient(lib.MqttClient)
	ttmqtt.Sub(lib.MqttClient, "#", temperatureReceivedHandler)
	dbName := "trop-tiede"
	// retrieves Mongo.Database instance
	lib.MyMusicAPIDB, lib.DBContext = db.InitDB(lib.Environment.MongoURL, lib.Environment.MongoPort, dbName)
	select {}
}
