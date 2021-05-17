package ormCustomerGroupEntityRelation

import (
	"errors"

	"github.com/chandantechpanion/adminpanel-service/pkg/customerGroupEntityRelation"
	"github.com/chandantechpanion/adminpanel-service/pkg/logger"

	"github.com/jinzhu/gorm"
)

type customerGroupEntityRelationRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

//NewCustomerGroupEntityRelationRepo ...
func NewCustomerGroupEntityRelationRepo(db *gorm.DB, log logger.LogInfoFormat) customerGroupEntityRelation.Repository {
	return &customerGroupEntityRelationRepo{
		db:  db,
		log: log,
	}
}

func (customer *customerGroupEntityRelationRepo) Save(c *customerGroupEntityRelation.CustomerGroupEntityRelation) error {

	// CName := customerGroupEntityRelation.CustomerGroupEntityRelation{}
	// err := customer.db.Where("customer_group_entity_relation_id = ? and user_info_userid =?", c., c.UserInfoUserID).First(&CName).Error

	// if err != nil {
	err := customer.db.Create(&c).Error
	if err != nil {
		customer.log.Errorf("error while creating the user, reason : %v", err)
		return err
	}
	// } else {
	// 	return errors.New("group name already exist")
	// }
	return nil
}

func (customer *customerGroupEntityRelationRepo) Get(id int64) (*customerGroupEntityRelation.CustomerGroupEntityRelation, error) {

	customerGroupEntityRelation := customerGroupEntityRelation.CustomerGroupEntityRelation{}
	err := customer.db.Where("id = ?", id).Find(&customerGroupEntityRelation).Error
	if err != nil {
		customer.log.Errorf("error while reading row, reason : %v", err)
		return nil, err
	}
	return &customerGroupEntityRelation, nil
}
func (customer *customerGroupEntityRelationRepo) GetAll() ([]*customerGroupEntityRelation.CustomerGroupEntityRelation, error) {

	customerMasterList := make([]*customerGroupEntityRelation.CustomerGroupEntityRelation, 0)
	err := customer.db.Find(&customerMasterList).Error
	if err != nil {
		return nil, err
	}
	return customerMasterList, nil
}

func (customer *customerGroupEntityRelationRepo) Delete(id int64) error {
	if customer.db.Delete(&customerGroupEntityRelation.CustomerGroupEntityRelation{}, "id = ?", id).Error != nil {
		return errors.New("error while deleting the customer")
	}
	return nil
}
