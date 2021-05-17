package ormCustomerMaster

import (
	"errors"

	"github.com/chandantechpanion/adminpanel-service/pkg/customerMaster"
	"github.com/chandantechpanion/adminpanel-service/pkg/logger"

	"github.com/jinzhu/gorm"
)

type customerMasterRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

//NewCustomerMasterRepo ...
func NewCustomerMasterRepo(db *gorm.DB, log logger.LogInfoFormat) customerMaster.Repository {
	return &customerMasterRepo{
		db:  db,
		log: log,
	}
}

func (i *customerMasterRepo) Save(c *customerMaster.CustomerMaster) error {

	CName := customerMaster.CustomerMaster{}

	err := i.db.Where("name = ?", c.Name).First(&CName).Error

	if err != nil {
		err := i.db.Create(&c).Error
		if err != nil {
			i.log.Errorf("error while creating the user, reason : %v", err)
			return err
		}
	} else {
		return errors.New("name already exist")
	}

	return nil
}
func (customer *customerMasterRepo) Get() ([]*customerMaster.CustomerMaster, error) {

	customerMasterList := make([]*customerMaster.CustomerMaster, 0)
	err := customer.db.Find(&customerMasterList).Error
	if err != nil {
		return nil, err
	}
	return customerMasterList, nil
}

func (customer *customerMasterRepo) GetByID(id int) (*customerMaster.CustomerMaster, error) {

	customerMaster := customerMaster.CustomerMaster{}
	err := customer.db.Where("id = ?", id).First(&customerMaster).Error
	if err != nil {
		return nil, err
	}
	return &customerMaster, nil
}

func (customer *customerMasterRepo) Delete(id int64) error {
	if customer.db.Delete(&customerMaster.CustomerMaster{}, "id = ?", id).Error != nil {
		return errors.New("error while deleting the customer")
	}
	return nil
}
