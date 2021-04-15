package main

import (
	"back/lib/utils"
	app "back/src"
)

func main(){
    serverEnv := utils.GetServerEnv()
    app.InitBackEnd(serverEnv)
}
