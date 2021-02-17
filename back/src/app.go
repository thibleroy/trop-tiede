package src

import (
	"back/lib"
	"back/lib/db"
	"back/lib/http"
	"back/lib/mqtt"
	"back/src/middlewares"
	"back/src/routes"
	"github.com/gorilla/mux"
)

func InitBackEnd(env lib.IEnvironment) {
	// saves environment in local variable
	lib.Environment = env

	// creates router and loads routes and methods handlers
	lib.Router = mux.NewRouter()
	routes.LoadRouters(lib.Router)

	// loads middlewares
	middlewares.LoadMiddlewares(lib.Router)
	dbName := "trop-tiede"

	// retrieves Mongo.Database instance
	lib.MyMusicAPIDB, lib.DBContext = db.InitDB(env.MongoURL, env.MongoPort, dbName)

	// initializes and connects the mqtt client
	lib.MqttClient = mqtt.InitMqttClient(env.MqttBrokerURL, env.MqttBrokerPort, env.MqttClientId, env.MqttUsername, env.MqttPassword)

	// initializes the http server with previously created Mux router
	http.InitWebServer(env.WebServerPort, lib.Router)
}

