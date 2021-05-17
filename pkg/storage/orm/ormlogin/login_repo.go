package ormlogin

import (
	"errors"

	"github.com/chandantechpanion/adminpanel-service/pkg/logger"
	"github.com/chandantechpanion/adminpanel-service/pkg/login"

	"github.com/jinzhu/gorm"
)

type loginRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

//NewloginRepo ...
func NewloginRepo(db *gorm.DB, log logger.LogInfoFormat) login.Repository {
	return &loginRepo{
		db:  db,
		log: log,
	}
}

func (l *loginRepo) SignIn(ln *login.Login) (string, error) {
	var login_data login.Login

	var UserId string
	row := l.db.Table("AdminUserInfo").Where("LoginUserName = ?", ln.LoginUserName).Select("UserId,LoginUserName,Password").Row() // (*sql.Row)
	err := row.Scan(&UserId, &login_data.LoginUserName, &login_data.Password)

	if err != nil {
		l.log.Errorf("signin failed with reason : %v", err)
		return "", err
	} else if login_data.LoginUserName == ln.LoginUserName && login_data.Password == ln.Password {
		return UserId, nil
	}
	return "", errors.New("please provide the correct password")
}
