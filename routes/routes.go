package routes

import (
	"github.com/gin-gonic/gin"
	"prototype/controllers/user_controller"
)

func InitRoutes(r *gin.Engine){
	userGroup := r.Group("/user_api")
	{
		userGroup.GET("/users", user_controller.Index)
		userGroup.GET("/users/:id", user_controller.Show)
		userGroup.POST("users", user_controller.Create)
		userGroup.POST("users/:id", user_controller.Update)
		userGroup.DELETE("users/:id", user_controller.Destroy)
	}
}