package adminUserInfo

import (
	"github.com/chandantechpanion/adminpanel-service/pkg/jwtService"
)

//Service interface
type Service interface {
	Get(login *Login) (string, error)
}

//ServiceAdminUserInfo ...
type ServiceAdminUserInfo struct {
	repo Repository
	jwt  jwtService.Service
}

//NewServiceforLogin to create new object
func NewServiceForAdminUserInfo(repo Repository, jwt jwtService.Service) Service {
	return &ServiceAdminUserInfo{
		repo: repo,
		jwt:  jwt,
	}
}

func (svc *ServiceAdminUserInfo) Get(login *Login) (string, error) {

	user, err := svc.repo.Get(login)
	if err == nil {
		token := svc.jwt.GenerateToken(user.UserID, true)
		return token, nil
	}
	return "", err
}
