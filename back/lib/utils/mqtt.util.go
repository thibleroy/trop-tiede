package utils

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	optionsReader := client.OptionsReader()
	fmt.Println("Connected to mqtt broker", optionsReader.Servers()[0])
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Println("Connect lost", err)
}

func wait(token mqtt.Token){
	token.Wait()
	if token.Error() != nil {
		panic(token.Error())
	}
}

func Sub(client mqtt.Client, topic string, cb mqtt.MessageHandler) {
	wait(client.Subscribe(topic, 0, cb))
	fmt.Println("Topic subscribed", topic)
}

func Pub(client mqtt.Client, topic string, payload string) {
	wait(client.Publish(topic, 0, false, payload))
	fmt.Println("Message published", )
}
func InitMqttClient(brokerAddr string, brokerPort int, clientId string, username string, password string) mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", brokerAddr, brokerPort))
	opts.SetClientID(clientId)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	return mqtt.NewClient(opts)
}

func ConnectMqttClient(client mqtt.Client){
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}


