package src

import (
	"back/lib"
	"back/lib/db"
	"back/lib/http"
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
	lib.MyMusicAPIDB = db.InitDB(env.MongoURL, env.MongoPort, lib.DBname)

	// initializes the http server with previously created Mux router
	http.InitWebServer(env.WebServerPort, lib.Router)
}

