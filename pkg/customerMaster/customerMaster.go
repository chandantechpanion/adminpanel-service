package customerMaster

import "time"

type CustomerMaster struct {
	ID         int64     `json:"id" db:"id" gorm:"not null;column:id;primary_key;AUTO_INCREMENT"`
	Name       string    `json:"name" db:"name" gorm:"not null;column:name"`
	CreateDate time.Time `json:"create_date" db:"create_date" gorm:"column:create_date;not null"`
}

func (CustomerMaster) TableName() string {
	return "CustomerMaster"
}
