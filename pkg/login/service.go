package login

import (
	"github.com/chandantechpanion/adminpanel-service/pkg/jwtService"
)

type Service interface {
	SignIn(userName string, password string) (string, error)
}

type loginService struct {
	repo Repository
	jwt  jwtService.Service
}

func NewLoginService(repo Repository, jwt jwtService.Service) Service {
	return &loginService{
		repo: repo,
		jwt:  jwt,
	}
}

func (svc *loginService) SignIn(email string, password string) (string, error) {

	userId, err := svc.repo.SignIn(&Login{LoginUserName: email, Password: password})
	if err == nil {
		token := svc.jwt.GenerateToken(userId, true)
		return token, nil
	}
	return "", err
}
