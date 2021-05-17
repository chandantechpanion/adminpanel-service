package adminUserInfo

//DBLoginInfo struct
type AdminUserInfo struct {
	UserID        string `json:"userid" db:"UserId" gorm:"column:UserId"`
	LoginUserName string `json:"login_user_name" db:"LoginUserName" gorm:"column:LoginUserName"`
	EmailID       string `json:"email_id" db:"EmailId" gorm:"column:EmailId"`
	PasswordHash  string `json:"password" db:"PasswordHash" gorm:"column:PasswordHash"`
	MobileNumber  string `json:"mobilenumber" db:"MobileNumber" gorm:"column:MobileNumber"`
	ActiveStatus  string `json:"activestatus" db:"ActiveStatus" gorm:"column:ActiveStatus"`
	Address1      string `json:"Address1" db:"Address1" gorm:"column:Address1"`
	Address2      string `json:"Address2" db:"Address2" gorm:"column:Address2"`
	City          string `json:"City" db:"City" gorm:"column:City"`
	State         string `json:"State" db:"State" gorm:"column:State"`
	PinCode       string `json:"Pincode" db:"Pincode" gorm:"column:Pincode"`
	Country       string `json:"Country" db:"Country" gorm:"column:Country"`
}

//TableName returns table name
func (AdminUserInfo) TableName() string {
	return "AdminUserInfo"
}
