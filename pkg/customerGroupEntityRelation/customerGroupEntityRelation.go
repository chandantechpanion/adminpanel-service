package customerGroupEntityRelation

import "time"

type CustomerGroupEntityRelation struct {
	ID               int64     `json:"id" db:"id" gorm:"not null;column:id;"`
	CustomerMasterID int64     `json:"customer_master_id" db:"customer_master_id" gorm:"not null;"`
	CustomerEntityID string    `json:"customer_entity_id" db:"customer_entity_id" gorm:"not null;column:customer_entity_id"`
	CustomerGroupID  int64     `json:"customer_group_id" db:"customer_group_id" gorm:"not null;"`
	CreateDate       time.Time `json:"create_date" db:"create_date" gorm:"column:create_date;not null"`
}

func (CustomerGroupEntityRelation) TableName() string {
	return "CustomerGroupEntityRelation"
}
