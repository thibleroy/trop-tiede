package src

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"back/src/lib"
	"back/src/lib/db"
	"back/src/lib/http"
	"back/src/lib/mqtt"
	"back/src/middlewares"
	"back/src/routes"
)

func InitBackEnd(env lib.IEnvironment) {
	// saves environment in local variable
	lib.Environment = env

	// prints running environment
	APIEnv, _ := json.Marshal(env)
	fmt.Println("API Environment", string(APIEnv))

	// creates router and loads routes and methods handlers
	lib.Router = mux.NewRouter()
	routes.LoadRouters(lib.Router)

	// loads middlewares
	middlewares.LoadMiddlewares(lib.Router)
	dbName := "MyMusicAPI"

	// retrieves Mongo.Database instance
	lib.MyMusicAPIDB, lib.DBContext = db.InitDB(env.MongoURL, env.MongoPort, dbName)

	// initializes the http server with previously created Mux router
	http.InitWebServer(env.WebServerPort, lib.Router)

	// initializes and connects the mqtt client
	lib.MqttClient = mqtt.InitMqttClient(env.MqttBrokerURL, env.MqttBrokerPort, env.MqttClientId, env.MqttUsername, env.MqttPassword)
}

