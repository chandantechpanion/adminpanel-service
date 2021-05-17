package customerGroup

import "time"

type CustomerGroup struct {
	ID          int64     `json:"id" db:"id" gorm:"not null;column:id;primary_key;AUTO_INCREMENT"`
	Name        string    `json:"name" db:"name" gorm:"not null;column:name"`
	CreatedDate time.Time `json:"created_date" db:"created_date" gorm:"column:created_date;not null"`
}

func (CustomerGroup) TableName() string {
	return "CustomerGroup"
}
