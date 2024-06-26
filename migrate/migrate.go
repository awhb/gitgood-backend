package main

import (
	"github.com/awhb/gitgood-backend/initialisers"
	"github.com/awhb/gitgood-backend/models"
)

func init() {
	initialisers.LoadEnvVariables()
	initialisers.ConnectToDB()
}

func main() {
	initialisers.DB.AutoMigrate(&models.Thread{})
}