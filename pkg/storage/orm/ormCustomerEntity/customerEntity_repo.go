package ormCustomerEntity

import (
	"errors"

	"github.com/chandantechpanion/adminpanel-service/pkg/customerEntity"
	"github.com/chandantechpanion/adminpanel-service/pkg/logger"

	"github.com/jinzhu/gorm"
)

type customerEntityRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

//NewCustomerEntityRepo ...
func NewCustomerEntityRepo(db *gorm.DB, log logger.LogInfoFormat) customerEntity.Repository {
	return &customerEntityRepo{
		db:  db,
		log: log,
	}
}

func (i *customerEntityRepo) Save(newEntity *customerEntity.CustomerEntity) error {

	entity := customerEntity.CustomerEntity{}
	err := i.db.Where("id = ? and customer_master_id = ?", newEntity.ID, newEntity.CustomerMasterID).First(&entity).Error
	if err != nil {
		err := i.db.Create(&newEntity).Error
		if err != nil {
			i.log.Errorf("error while creating the entity, reason : %v", err)
			return err
		}
	} else {
		return errors.New("entity is already exist")
	}

	return nil
}

func (customer *customerEntityRepo) GetAll() ([]*customerEntity.CustomerEntity, error) {

	customerEntity := make([]*customerEntity.CustomerEntity, 0)
	err := customer.db.Find(&customerEntity).Error
	if err != nil {
		customer.log.Errorf("error while reading row, reason : %v", err)
		return nil, err
	}
	return customerEntity, nil
}

func (customer *customerEntityRepo) Get(id int64) ([]*customerEntity.CustomerEntity, error) {

	customerEntity := make([]*customerEntity.CustomerEntity, 0)
	err := customer.db.Where("customer_master_id = ?", id).Find(&customerEntity).Error
	if err != nil {
		return nil, err
	}
	return customerEntity, nil
}
