package customerEntity

import "time"

type CustomerEntity struct {
	ID               string    `json:"id" db:"id" gorm:"not null;column:id;"`
	CustomerMasterID int64     `json:"customer_master_id" db:"customer_master_id" gorm:"not null;column:customer_master_id;"`
	Name             string    `json:"name" db:"name" gorm:"not null;column:name"`
	CreatedDate      time.Time `json:"created_date" db:"created_date" gorm:"column:created_date;not null"`
}

func (CustomerEntity) TableName() string {
	return "CustomerEntity"
}
