package main

import (
	"elix/server"
	// "elix/socket"
	"elix/utils"
)

func main() {
	utils.StartUpScreen()
	utils.FSAssureStoreExist("elix")
	serverPort := utils.ServerPort()
	go server.Server(serverPort)
	select {}
}
