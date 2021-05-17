package buyerUserInfo

import "github.com/chandantechpanion/adminpanel-service/pkg/helper"

type Service interface {
	Save(user *CreateBuyerUserInfo) error
	Get(GST string) (*BuyerUserInfo, error)
}
type buyerUserInfoService struct {
	repo Repository
}

func NewBuyerUserInfoService(repo Repository) Service {
	return &buyerUserInfoService{
		repo: repo,
	}
}

func (svc *buyerUserInfoService) Save(user *CreateBuyerUserInfo) error {
	buyerUser := BuyerUserInfo{}
	helper.DeepCopy(user, buyerUser)
	return svc.repo.Save(&buyerUser)
}

func (svc *buyerUserInfoService) Get(GST string) (*BuyerUserInfo, error) {
	return svc.repo.Get(GST)
}
