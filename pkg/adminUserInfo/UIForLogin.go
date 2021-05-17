package adminUserInfo

//UIForLogin is a struct of inputs received from user
type Login struct {
	LoginUserName string `json:"login_user_name" db:"LoginUserName" gorm:"column:LoginUserName"`
	Password      string `json:"password" db:"Password" gorm:"column:Password"`
}
