package routes

import (
	"github.com/awhb/gitgood-backend/controllers"
	"github.com/awhb/gitgood-backend/middleware"
	"github.com/gin-gonic/gin"
)

func Threads(route *gin.RouterGroup) {
	threads := route.Group("/threads")
	{
		threads.POST("", middleware.RequireAuth, controllers.ThreadsCreate)
		threads.GET("", controllers.ThreadsIndex)
		threads.GET("/:id", controllers.ThreadsShow)
		threads.PUT("/:id", controllers.ThreadsUpdate)
		threads.DELETE("/:id", controllers.ThreadsDelete)
	}
}
