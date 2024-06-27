package routes

import (
	"github.com/awhb/gitgood-backend/controllers"
	"github.com/awhb/gitgood-backend/middleware"
	"github.com/gin-gonic/gin"
)

func Comments(route *gin.RouterGroup) {
	comments := route.Group("")
	{
		comments.POST("/comments/", middleware.RequireAuth, controllers.CommentsCreate)
		comments.GET("/threads/:thread_id/comments", controllers.CommentsIndex)
		comments.PUT("/comments/:id", controllers.CommentsUpdate)
		comments.DELETE("/comments/:id", controllers.CommentsDelete)
	}
}