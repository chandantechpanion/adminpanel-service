package login

type Login struct {
	LoginUserName string `json:"loginUserName" db:"LoginUserName" gorm:"not null"`
	Password      string `json:"password" db:"Password" gorm:"not null"`
}
