package customerMember

import "time"

type CustomerMember struct {
	ID                            int64     `json:"id" db:"id" gorm:"not null;column:id;"`
	CustomerGroupEntityRelationID int64     `json:"customer_group_entity_relation_id" db:"customer_group_entity_relation_id" gorm:"not null;"`
	UserInfoUserID                string    `json:"user_info_userid" db:"user_info_userid" gorm:"not null;column:user_info_userid"`
	CreateDate                    time.Time `json:"create_date" db:"create_date" gorm:"column:create_date;not null"`
}

func (CustomerMember) TableName() string {
	return "CustomerMember"
}
