package services

import (
	"go-blog/app/models/user"
	"go-blog/app/repositories"
	"go-blog/app/requests"
	"go-blog/pkg/auth"
	"go-blog/pkg/jwt"
)

type AuthService interface {
	HandleLogin(data *requests.LoginRequest) (user.User, error)
	HandleLogout()
	HandleUpdatePassword()
}

type authService struct {
	userRepository repositories.UserRepository
}

func NewAuthService() AuthService {
	return &authService{
		userRepository: repositories.NewUserRepository(),
	}
}

func (s *authService) HandleLogin(data *requests.LoginRequest) (user.User, error) {
	userModel, err := auth.Attempt(data.Username, data.Password)
	if err != nil {
		return userModel, err
	}
	userModel.Token = jwt.NewJWT().IssueToken(userModel.ID, userModel.Nickname)
	return userModel, err
}

func (s *authService) HandleLogout() {

}

func (s *authService) HandleUpdatePassword() {

}
