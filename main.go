package main

import (
    "gossip-forum-backend/database"
    "gossip-forum-backend/initialisers"
    "gossip-forum-backend/routes"


    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func init() {
    initialisers.LoadEnvVariables()
}

func main() {
    r := gin.Default()
    r.Use(cors.Default())  // Allow CORS

    database.ConnectDatabase()

    v1 := r.Group("/api/v1")
    {
        routes.Comments(v1)
        routes.Threads(v1)
        routes.Users(v1)
    }
  
    r.Run() // listen and serve on 0.0.0.0:8080
}
