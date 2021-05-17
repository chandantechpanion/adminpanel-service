package ormBuyerUserInfo

import (
	"errors"

	"github.com/chandantechpanion/adminpanel-service/pkg/buyerUserInfo"
	"github.com/chandantechpanion/adminpanel-service/pkg/logger"

	"github.com/jinzhu/gorm"
)

type buyerUserInfoRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

func NewBuyerUserInfoRepo(db *gorm.DB, log logger.LogInfoFormat) buyerUserInfo.Repository {
	return &buyerUserInfoRepo{
		db:  db,
		log: log,
	}
}

func (b *buyerUserInfoRepo) Save(user *buyerUserInfo.BuyerUserInfo) error {
	b.log.Debugf("creating the user with email : %v", user.EmailId)

	findUser := buyerUserInfo.BuyerUserInfo{}
	err := b.db.Where("EmailId = ? or MobileNumber = ?", user.EmailId, user.MobileNumber).First(&findUser).Error

	if err != nil {
		err := b.db.Create(&user).Error
		if err != nil {
			b.log.Errorf("error while creating the user, reason : %v", err)
			return err
		}
	} else {
		return errors.New("user is already exist")
	}
	return nil
}

func (b *buyerUserInfoRepo) Get(GST string) (*buyerUserInfo.BuyerUserInfo, error) {

	b.log.Debugf("Get the user by GST : %v", GST)
	user := buyerUserInfo.BuyerUserInfo{}
	err := b.db.Where("GST = ?", GST).Find(&user).Error
	if err != nil {
		b.log.Errorf("error while get the user, reason : %v", err)
		return nil, err
	}
	return &user, nil
}
