package ormCustomerMember

import (
	"errors"

	"github.com/chandantechpanion/adminpanel-service/pkg/customerMember"
	"github.com/chandantechpanion/adminpanel-service/pkg/logger"

	"github.com/jinzhu/gorm"
)

type customerMemberRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

//NewCustomerMemberRepo ...
func NewCustomerMemberRepo(db *gorm.DB, log logger.LogInfoFormat) customerMember.Repository {
	return &customerMemberRepo{
		db:  db,
		log: log,
	}
}

func (customer *customerMemberRepo) Save(c *customerMember.CustomerMember) error {

	CName := customerMember.CustomerMember{}

	// row := customer.db.Table("CustomerMember").Where("customer_group_entity_relation_id = ? and user_info_userid =?", c.CustomerGroupEntityRelationID, c.UserInfoUserID).Select("customer_group_entity_relation_id,user_info_userid").Row()
	// err := row.Scan(&CName.CustomerGroupEntityRelationID, &CName.UserInfoUserID)
	err := customer.db.Where("customer_group_entity_relation_id = ? and user_info_userid =?", c.CustomerGroupEntityRelationID, c.UserInfoUserID).First(&CName).Error

	if err != nil {
		err := customer.db.Create(&c).Error
		if err != nil {
			customer.log.Errorf("error while creating the user, reason : %v", err)
			return err
		}
	} else {
		return errors.New("group name already exist")
	}
	return nil
}

func (customer *customerMemberRepo) Get(id int64) (*customerMember.CustomerMember, error) {

	customerMember := customerMember.CustomerMember{}
	err := customer.db.Where("id = ?", id).Find(&customerMember).Error
	if err != nil {
		customer.log.Errorf("error while reading row, reason : %v", err)
		return nil, err
	}
	return &customerMember, nil
}

func (customer *customerMemberRepo) Delete(id int64) error {
	if customer.db.Delete(&customerMember.CustomerMember{}, "id = ?", id).Error != nil {
		return errors.New("error while deleting the customer")
	}
	return nil
}
