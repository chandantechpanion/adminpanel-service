package buyerUserInfo

type Repository interface {
	Save(user *BuyerUserInfo) error
	Get(GST string) (*BuyerUserInfo, error)
}
