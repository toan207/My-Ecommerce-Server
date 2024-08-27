package controller

import (
	"github.com/gin-gonic/gin"
	"my-ecom.com/api/internal/service"
	"my-ecom.com/api/pkg/response"
)

type UserController struct {
	userService *service.UserService
}

func IUserController() *UserController {
	return &UserController{
		userService: service.IUserService(),
	}
}

// controller -> service -> repo -> models -> dbs
func (uc *UserController) Pong(c *gin.Context) {
	response.SuccessResponse(c, 20001, []string{"Toan", "Toan1"})
}
