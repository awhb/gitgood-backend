package routes

import (
	"github.com/awhb/gitgood-backend/controllers"
	"github.com/awhb/gitgood-backend/middleware"
	"github.com/gin-gonic/gin"
)

func Comments(route *gin.RouterGroup) {
	comments := route.Group("/comments")
	{
		comments.POST("/", middleware.RequireAuth, controllers.CommentsCreate)
		comments.GET("/:thread_id", controllers.CommentsIndex)
		comments.PUT("/:id", controllers.CommentsUpdate)
		comments.DELETE("/:id", controllers.CommentsDelete)
	}
}