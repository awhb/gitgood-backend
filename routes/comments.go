package routes

import (
	"gossip-backend/controllers"
	"github.com/gin-gonic/gin"
)

func Comments(route *gin.RouterGroup) {
		comments := route.Group("/comments")
		{
				comments.POST("/comments", controllers.CreateComment)
				comments.GET("/threads/:thread_id/comments", controllers.GetComments)
				comments.PUT("/comments/:id", controllers.UpdateComment)
				comments.DELETE("/comments/:id", controllers.DeleteComment)
		}
}