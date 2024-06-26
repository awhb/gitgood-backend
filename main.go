package main

import (
    "github.com/awhb/gossip-backend/initialisers"
    "github.com/awhb/gossip-backend/routes"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func init() {
    initialisers.LoadEnvVariables()
    initialisers.ConnectToDB()
}

func main() {
    r := gin.Default()
    r.Use(cors.Default())  // Allow CORS

    v1 := r.Group("/api/v1")
    {
        routes.Comments(v1)
        routes.Threads(v1)
        routes.Users(v1)
    }

    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "data": "Hello, World!",
        })
    })
  
    r.Run() // listen and serve on 0.0.0.0:8080
}
