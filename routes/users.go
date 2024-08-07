package routes

import (
	"github.com/awhb/gitgood-backend/controllers"
	"github.com/awhb/gitgood-backend/middleware"
	"github.com/gin-gonic/gin"
)

func Users(route *gin.RouterGroup) {
	users := route.Group("")
	{
		users.POST("/register", controllers.Register)
		users.POST("/login", controllers.Login)
		users.GET("/validate", middleware.RequireAuth, controllers.Validate)
	}
}
