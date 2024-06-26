package main

import (
	"github.com/awhb/gossip-backend/initialisers"
	"github.com/awhb/gossip-backend/models"
)

func init() {
	initialisers.LoadEnvVariables()
	initialisers.ConnectToDB()
}

func main() {
	initialisers.DB.AutoMigrate(&models.Thread{})
}