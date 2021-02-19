package main

import (
	"back/lib"
	ttmqtt "back/lib/mqtt"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var temperatureReceivedHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received 1 message: %s from topic: %s\n", msg.Payload(), msg.Topic())

}

func main(){
	lib.Environment = lib.GetServerEnv()
	lib.MqttClient = ttmqtt.InitMqttClient(lib.Environment.MqttBrokerURL, lib.Environment.MqttBrokerPort, lib.Environment.MqttClientId, lib.Environment.MqttUsername, lib.Environment.MqttPassword)
	ttmqtt.ConnectMqttClient(lib.MqttClient)
	ttmqtt.Sub(lib.MqttClient, lib.Environment.MqttTemperatureTopic, temperatureReceivedHandler)
	for {
	}
}
