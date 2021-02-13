package lib

import (
	"context"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var MyMusicAPIDB mongo.Database
var DBContext context.Context
var Router *mux.Router
var Environment IEnvironment
var MqttClient mqtt.Client
