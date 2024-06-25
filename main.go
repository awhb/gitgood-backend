package main

import (
    "gossip-forum-backend/controllers"
    "gossip-forum-backend/database"
    "gossip-forum-backend/routes"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.Use(cors.Default())  // Allow CORS

    database.ConnectDatabase()
    routes.SetupRoutes(r)

    r.Run() // listen and serve on 0.0.0.0:8080
}
