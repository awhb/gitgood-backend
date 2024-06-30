package main

import (
    "github.com/awhb/gitgood-backend/initialisers"
	"github.com/awhb/gitgood-backend/routes"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func init() {
	// Disable in production
    initialisers.LoadEnvVariables()
	
    initialisers.ConnectToDB()
	initialisers.SyncDB()
}

func main() {
	// gin.SetMode(gin.ReleaseMode)

    r := gin.Default()

    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"}, // Update with your frontend URL
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    api := r.Group("")
    {
        routes.Comments(api)
        routes.Threads(api)
        routes.Users(api)
    }
  
    r.Run() // listen and serve on 0.0.0.0:8080
}

