package routes

import (
	"web-forum-backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
			api.POST("/register", controllers.Register)
			api.POST("/login", controllers.Login)

			api.POST("/threads", controllers.CreateThread)
			api.GET("/threads", controllers.GetThreads)
			api.GET("/threads/:id", controllers.GetThread)
			api.PUT("/threads/:id", controllers.UpdateThread)
			api.DELETE("/threads/:id", controllers.DeleteThread)
	}

	return r
}