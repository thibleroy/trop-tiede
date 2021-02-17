package main

import (
	"back/lib"
	app "back/src"
)

func main(){
    serverEnv := lib.GetServerEnv()
    app.InitBackEnd(serverEnv)
}
