package main

import (
	"back/lib"
	ttmqtt "back/lib/mqtt"
	"encoding/json"
	"fmt"
	"time"
)

func main(){
	lib.Environment = lib.GetServerEnv()
	lib.MqttClient = ttmqtt.InitMqttClient(lib.Environment.MqttBrokerURL, lib.Environment.MqttBrokerPort, lib.Environment.MqttClientId, lib.Environment.MqttUsername, lib.Environment.MqttPassword)
	ttmqtt.ConnectMqttClient(lib.MqttClient)
	for {
		data := lib.IRoomData{
			Temperature: 20.5,
			Device:      lib.IDevice{
				DeviceDescription: lib.IDeviceDescription{
					IDescription: lib.IDescription{
						Name:    "ESP32 uPesy",
						Details: "La carte uPesy ESP32 WROVER DevKit est basée sur un ESP32. " +
							"Cette carte peut être mise sur une breadboard facilement car les 2 " +
							"côtés de la carte sont accessibles pour mettre des fils sur la breadboard.",
					},
					SerialNumber: "azhefdzj12frezd",
				},
			},
		}
		payload,_ := json.Marshal(data)
		fmt.Println("*******************")
		fmt.Printf("%+v\n", data)
		fmt.Println("*******************")
		ttmqtt.Pub(lib.MqttClient, lib.Environment.MqttTemperatureTopic, string(payload))
		time.Sleep(5 * time.Second)
	}
}
