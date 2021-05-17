package adminUserInfo

//Repository interface
type Repository interface {
	Get(login *Login) (*AdminUserInfo, error)
}
