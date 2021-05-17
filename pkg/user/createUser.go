package user

type CreateUser struct {
	LoginUserName string `json:"login_user_name" db:"LoginUserName" gorm:"varchar(50);column:LoginUserName;not null"`
	FirstName     string `json:"first_name" db:"FirstName" gorm:"varchar(50);column:FirstName;not null"`
	LastName      string `json:"last_name" db:"LastName" gorm:"varchar(50);column:LastName"`
	EmailId       string `json:"email" db:"EmailId" gorm:"varchar(200);not null;unique"`
	PhoneNumber   string `json:"phone_number" db:"PhoneNumber" gorm:"not null;column:PhoneNumber"`
	Password      string `json:"password" db:"password" gorm:"varchar(50);not null"`
}
