package routes

import (
	"github.com/awhb/gossip-backend/controllers"
	"github.com/gin-gonic/gin"
)

func Users(route *gin.RouterGroup) {
	users := route.Group("")
	{
		users.POST("/register", controllers.Register)
		users.POST("/login", controllers.Login)
	}
}
