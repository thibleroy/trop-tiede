package src

import (
	"back/lib"
	"back/lib/utils"
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

	// retrieves Mongo.Database instance
	lib.MyMusicAPIDB = utils.InitDB(env.MongoURL, env.MongoPort, lib.DBname)

	// initializes the http server with previously created Mux router
	utils.InitWebServer(env.WebServerPort, lib.Router)
}

