package main

import (
	"back/lib"
	ttmqtt "back/lib/mqtt"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strings"
)

var tempStore []lib.IRoomData

var temperatureReceivedHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Println(strings.Split(msg.Topic(),"/"))
	room := strings.Split(msg.Topic(),"/")[1]
	dataType := strings.Split(msg.Topic(),"/")[2]
	fmt.Println("room", room)
	fmt.Println("datatype", dataType)
	fmt.Printf("Received 1 message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	var currentTemp lib.IRoomData
	err := json.Unmarshal(msg.Payload(), &currentTemp)
	if err != nil {
		panic(err)
	}
	tempStore = append(tempStore, currentTemp)
	fmt.Println("New store", tempStore)
}

func main(){
	lib.Environment = lib.GetServerEnv()
	lib.MqttClient = ttmqtt.InitMqttClient(lib.Environment.MqttBrokerURL, lib.Environment.MqttBrokerPort, lib.Environment.MqttClientId, lib.Environment.MqttUsername, lib.Environment.MqttPassword)
	ttmqtt.ConnectMqttClient(lib.MqttClient)
	ttmqtt.Sub(lib.MqttClient, lib.Environment.MqttTemperatureTopic, temperatureReceivedHandler)
	select {}
}
