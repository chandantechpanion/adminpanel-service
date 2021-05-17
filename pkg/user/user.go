package user

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	UserId           string    `json:"user_id" db:"UserId" gorm:"column:UserId;primary_key;type:uuid"`
	LoginUserName    string    `json:"login_user_name" db:"LoginUserName" gorm:"varchar(50);column:LoginUserName;not null"`
	FirstName        string    `json:"first_name" db:"FirstName" gorm:"varchar(50);column:FirstName;not null"`
	LastName         string    `json:"last_name" db:"LastName" gorm:"varchar(50);column:LastName"`
	EmailId          string    `json:"email" db:"EmailId" gorm:"varchar(200);not null;unique"`
	PhoneNumber      string    `json:"phone_number" db:"PhoneNumber" gorm:"not null;column:PhoneNumber"`
	Password         string    `json:"password" db:"password" gorm:"varchar(50);not null"`
	IsEmailConfirmed bool      `json:"is_email_confirmed" db:"IsEmailConfirmed" gorm:"not null"`
	ActiveStatus     bool      `json:"active_status" db:"ActiveStatus"`
	RegistrationDate time.Time `json:"registration_date" db:"RegistrationDate" gorm:"column:RegistrationDate"`
}

func (User) TableName() string {
	return "UserInfo"
}

func (u *User) BeforeCreate(scope *gorm.Scope) {
	uuid := uuid.NewV4()
	scope.SetColumn("ID", uuid.String())
}
