package main

import (
	"back/lib"
	ttmqtt "back/lib/mqtt"
	"encoding/json"
	"time"
)

func main(){
	lib.Environment = lib.GetServerEnv()
	lib.MqttClient = ttmqtt.InitMqttClient(lib.Environment.MqttBrokerURL, lib.Environment.MqttBrokerPort, lib.Environment.MqttClientId, lib.Environment.MqttUsername, lib.Environment.MqttPassword)
	ttmqtt.ConnectMqttClient(lib.MqttClient)
	for {
		payload,_ := json.Marshal(lib.IRoomData{
			Temperature: 20.5,
			Time: time.Now(),
		})
		ttmqtt.Pub(lib.MqttClient, lib.Environment.MqttTemperatureTopic, string(payload))
		time.Sleep(10 * time.Second)
	}
}
