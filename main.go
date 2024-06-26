package main

import (
    "github.com/awhb/gitgood-backend/initialisers"
    "github.com/awhb/gitgood-backend/routes"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func init() {
    initialisers.LoadEnvVariables()
    initialisers.ConnectToDB()
}

func main() {
	gin.SetMode(gin.ReleaseMode)

    r := gin.Default()
    r.Use(cors.Default())  // Allow CORS

    v1 := r.Group("/api/v1")
    {
        routes.Comments(v1)
        routes.Threads(v1)
        routes.Users(v1)
    }
  
    r.Run() // listen and serve on 0.0.0.0:8080
}
