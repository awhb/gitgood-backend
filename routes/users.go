package routes

import (
	"github.com/awhb/gitgood-backend/controllers"
	"github.com/gin-gonic/gin"
)

func Users(route *gin.RouterGroup) {
	users := route.Group("")
	{
		users.POST("/signup", controllers.Signup)
		users.POST("/login", controllers.Login)
		users.GET("/validate", controllers.Validate)
	}
}
