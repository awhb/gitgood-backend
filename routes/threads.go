package routes

import (
	"gossip-backend/controllers"
	"github.com/gin-gonic/gin"
)

func Threads(route *gin.RouterGroup) {
		threads := route.Group("/threads")
		{
				threads.POST("/threads", controllers.CreateThread)
				threads.GET("/threads", controllers.GetThreads)
				threads.GET("/threads/:id", controllers.GetThread)
				threads.PUT("/threads/:id", controllers.UpdateThread)
				threads.DELETE("/threads/:id", controllers.DeleteThread)
		}
}
