package routes

import (
	"gossip-backend/controllers"
	"github.com/gin-gonic/gin"
)

func Users(route *gin.RouterGroup) {
		users := route.Group("/users")
		{
				users.POST("/register", controllers.Register)
				users.POST("/login", controllers.Login)
		}
}
