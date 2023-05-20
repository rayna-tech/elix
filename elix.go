package main

import (
	"elix/server"
	"os"
	// "elix/socket"
	"elix/utils"

	"github.com/joho/godotenv"
)

func main() {
	// Basics
	utils.StartUpScreen()
	utils.FSAssureStoreExist("elix")

	// ENV
	err := godotenv.Load()
	if err != nil {
		utils.Logger.Error("Failed to load .env file.")
		return
	}

	authToken := os.Getenv("AUTH_TOKEN")
	if authToken == "" {
		utils.Logger.Error("\"AUTH_TOKEN\n value required in .env file.")
		return
	}

	// Server
	serverPort := utils.ServerPort()
	go server.Server(serverPort)
	select {}
}
