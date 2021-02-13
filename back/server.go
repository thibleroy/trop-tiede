package main

import (
    app "back/src"
    "back/src/lib"
)

func main(){
    serverEnv := lib.GetServerEnv()
    app.InitBackEnd(serverEnv)
}
