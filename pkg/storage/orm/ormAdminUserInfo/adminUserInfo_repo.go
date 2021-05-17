package ormCustomerEntity

import (
	"github.com/chandantechpanion/adminpanel-service/pkg/adminUserInfo"
	"github.com/chandantechpanion/adminpanel-service/pkg/logger"

	"github.com/jinzhu/gorm"
)

type adminUserInfoRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

//NewAdminUserInfoRepo ...
func NewAdminUserInfoRepo(db *gorm.DB, log logger.LogInfoFormat) adminUserInfo.Repository {
	return &adminUserInfoRepo{
		db:  db,
		log: log,
	}
}

func (l *adminUserInfoRepo) Get(login *adminUserInfo.Login) (*adminUserInfo.AdminUserInfo, error) {

	User := adminUserInfo.AdminUserInfo{}
	err := l.db.Where("LoginUserName = ?", login.LoginUserName).First(&User).Error
	if err != nil {
		l.log.Errorf("error while reading row, reason : %v", err)
		return nil, err
	}
	return &User, nil
}
