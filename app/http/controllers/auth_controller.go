package controllers

import (
	"go-blog/app/requests"
	"go-blog/app/services"
	"go-blog/pkg/request"
	"go-blog/pkg/response"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		authService: services.NewAuthService(),
	}
}

func (ac *AuthController) UpdatePassword(c *gin.Context) {

}

func (ac *AuthController) Login(c *gin.Context) {
	data := &requests.LoginRequest{}
	if ok := request.Validate(c, data); !ok {
		return
	}
	userModel, err := ac.authService.HandleLogin(data)
	if err != nil {
		response.Error(c, err)
	} else {
		response.Data(c, gin.H{
			"token": userModel.Token,
		})
	}
}

func (ac *AuthController) Logout(c *gin.Context) {

}
