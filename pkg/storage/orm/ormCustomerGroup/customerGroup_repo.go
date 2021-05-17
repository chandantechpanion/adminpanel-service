package ormCustomerGroup

import (
	"errors"

	"github.com/chandantechpanion/adminpanel-service/pkg/customerGroup"
	"github.com/chandantechpanion/adminpanel-service/pkg/logger"

	"github.com/jinzhu/gorm"
)

type customerGroupRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

//NewCustomerGroupRepo ...
func NewCustomerGroupRepo(db *gorm.DB, log logger.LogInfoFormat) customerGroup.Repository {
	return &customerGroupRepo{
		db:  db,
		log: log,
	}
}

func (i *customerGroupRepo) Save(c *customerGroup.CustomerGroup) error {

	CName := customerGroup.CustomerGroup{}

	err := i.db.Where("name = ?", c.Name).First(&CName).Error

	if err != nil {
		err := i.db.Create(&c).Error
		if err != nil {
			i.log.Errorf("error while creating the user, reason : %v", err)
			return err
		}
	} else {
		return errors.New("group name already exist")
	}

	return nil
}

// func (customer *customerGroupRepo) Get(id int64) (*customerGroup.CustomerGroup, error) {

// 	customerGroup := customerGroup.CustomerGroup{}
// 	err := customer.db.Where("id = ?", id).Find(&customerGroup).Error
// 	if err != nil {
// 		customer.log.Errorf("error while reading row, reason : %v", err)
// 		return nil, err
// 	}
// 	return &customerGroup, nil
// }

func (customer *customerGroupRepo) Get(userId int64, OffSet, FetchNext int) ([]*customerGroup.CustomerGroup, error) {
	customer.log.Debugf("id : %v", userId)

	customerList := make([]*customerGroup.CustomerGroup, 0)
	err := customer.db.Order("id desc").Offset(OffSet).Limit(FetchNext).Find(&customerList).Error
	if err != nil {
		return nil, err
	}
	return customerList, nil
}

func (customer *customerGroupRepo) Get_Pagination(userId int64) (int, error) {
	// customer.log.Debugf("id - Get_Pagination : %v", userId)
	var resultCount int
	err := customer.db.Model(&customerGroup.CustomerGroup{}).Count(&resultCount).Error
	if err != nil {
		return 0, err
	}
	return resultCount, nil
}

func (customer *customerGroupRepo) Delete(id int64) error {
	if customer.db.Delete(&customerGroup.CustomerGroup{}, "id = ?", id).Error != nil {
		return errors.New("error while deleting the customer")
	}
	return nil
}
