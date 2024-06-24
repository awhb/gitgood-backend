package main

import (
    "gossip-forum-backend/controllers"
    "gossip-forum-backend/database"
    // "gossip-forum-backend/routes"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Allow CORS
    r.Use(cors.Default())

    // Connect to database
    database.ConnectDatabase()

    // Set up routes
    // routes.SetupRoutes(r)
		r.POST("/login", controllers.Login)
		r.POST("/register", controllers.Register)

    r.Run() // listen and serve on 0.0.0.0:8080
}
