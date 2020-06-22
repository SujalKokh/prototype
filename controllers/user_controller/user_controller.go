package user_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	userService "prototype/services/user_service"
)

func Index(c *gin.Context){	
	// u.GetAllUsers()
	// userServiceThing := &UserService{}
	// userServiceThing.GetAllUsers()
	userService.GetAllUsers()
	c.JSON(http.StatusOK, "Index/All")
}

func Show(c *gin.Context){
	c.JSON(http.StatusOK, "Show Action")
}

func Create(c *gin.Context){
	c.JSON(http.StatusOK, "Create Action")
}

func Update(c *gin.Context){
	c.JSON(http.StatusOK, "Update Action")
}

func Destroy(c *gin.Context){
	c.JSON(http.StatusOK, "Destroy action")
}
