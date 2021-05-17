package ormUser

import (
	"errors"
	"fmt"

	"github.com/chandantechpanion/adminpanel-service/pkg/logger"
	"github.com/chandantechpanion/adminpanel-service/pkg/user"

	"github.com/jinzhu/gorm"
)

type userRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

func NewUserRepo(db *gorm.DB, log logger.LogInfoFormat) user.Repository {
	return &userRepo{db, log}
}

func (u *userRepo) Delete(id string) error {
	u.log.Debugf("deleting the user with id : %s", id)

	if u.db.Delete(&user.User{}, "user_id = ?", id).Error != nil {
		errMsg := fmt.Sprintf("error while deleting the user with id : %s", id)
		u.log.Errorf(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func (u *userRepo) GetAll() ([]*user.User, error) {
	u.log.Debug("get all the users")

	users := make([]*user.User, 0)
	err := u.db.Find(&users).Error
	if err != nil {
		u.log.Debug("no single user found")
		return nil, err
	}
	return users, nil
}

func (u *userRepo) GetByID(id string) (*user.User, error) {
	u.log.Debugf("get user details by id : %s", id)

	user := &user.User{}
	err := u.db.Where("UserId = ?", id).First(&user).Error
	user.Password = ""
	if err != nil {
		u.log.Errorf("user not found with id : %s, reason : %v", id, err)
		return nil, err
	}
	return user, nil
}

func (u *userRepo) Store(usr *user.User) error {
	u.log.Debugf("creating the user with email : %v", usr.EmailId)

	findUser := user.User{}
	err := u.db.Where("EmailId = ?", usr.EmailId).First(&findUser).Error

	if err != nil {
		err := u.db.Create(&usr).Error
		if err != nil {
			u.log.Errorf("error while creating the user, reason : %v", err)
			return err
		}
	} else {
		return errors.New("user is already exist")
	}
	return nil
}

func (u *userRepo) Update(usr *user.User) error {
	u.log.Debugf("updating the user, user_id : %v", usr.UserId)

	err := u.db.Model(&usr).Updates(user.User{FirstName: usr.FirstName, LastName: usr.LastName, Password: usr.Password, PhoneNumber: usr.PhoneNumber}).Error
	if err != nil {
		u.log.Errorf("error while updating the user, reason : %v", err)
		return err
	}
	return nil
}
