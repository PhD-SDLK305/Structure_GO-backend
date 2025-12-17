package controllers

import (
	"gobackend/iternal/services"
	"gobackend/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

// receiver struct
func (uc *UserController) GetUserById(c *gin.Context) {
	// response.SuccessResponse(c, 20001, []string{"quan"})
	response.ErrorResponse(c, 20003, "")
}
