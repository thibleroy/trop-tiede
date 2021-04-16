package lib

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var MyDB mongo.Database
var Router *mux.Router
var Environment IEnvironment
var MqttClient mqtt.Client
