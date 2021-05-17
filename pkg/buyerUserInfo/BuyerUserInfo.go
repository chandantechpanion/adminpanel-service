package buyerUserInfo

type BuyerUserInfo struct {
	UserId       string `json:"user_id" db:"UserId" gorm:"column:UserId;primary_key;type:uuid"`
	FirstName    string `json:"first_name" db:"FirstName" gorm:"column:FirstName;"`
	LastName     string `json:"last_name" db:"LastName" gorm:"column:LastName;"`
	EmailId      string `json:"email_id" db:"EmailId" gorm:"column:EmailId;"`
	MobileNumber string `json:"mobile_number" db:"MobileNumber" gorm:"column:mobile_number;"`
	TradeName    string `json:"TradeName" db:"TradeName" gorm:"column:TradeName;"`
	LegalName    string `json:"LegalName" db:"LegalName" gorm:"column:LegalName;"`
	CustomerGST  string `json:"CustomerGST" db:"CustomerGST" gorm:"column:CustomerGST;"`
	StateCode    string `json:"StateCode" db:"StateCode" gorm:"column:StateCode;"`
	Address1     string `json:"Address1" db:"Address1" gorm:"column:Address1;"`
	Address2     string `json:"Address2" db:"Address2" gorm:"column:Address2;"`
	City         string `json:"City" db:"City" gorm:"column:City;"`
	State        string `json:"State" db:"State" gorm:"column:State;"`
	Pincode      string `json:"Pincode" db:"Pincode" gorm:"column:Pincode;"`
	Country      string `json:"Country" db:"Country" gorm:"column:Country;"`
	PAN          string `json:"PAN" db:"PAN" gorm:"column:PAN;"`
	GST          string `json:"GST" db:"GST" gorm:"column:GST;"`
	CIN          string `json:"CIN" db:"CIN" gorm:"column:CIN;"`
}

func (BuyerUserInfo) TableName() string {
	return "BuyerUserInfo"
}
