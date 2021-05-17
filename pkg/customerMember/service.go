package customerMember

//Service ...
type Service interface {
	Get(id int64) (*CustomerMember, error)
	Save(custMember *CustomerMember) error
	Delete(id int64) error
}

type customerMemberService struct {
	repo Repository
}

//NewCustomerMemberService ...
func NewCustomerMemberService(repo Repository) Service {
	return &customerMemberService{
		repo: repo,
	}
}

func (svc *customerMemberService) Get(id int64) (*CustomerMember, error) {
	return svc.repo.Get(id)
}

func (svc *customerMemberService) Save(custMember *CustomerMember) error {
	return svc.repo.Save(custMember)
}

func (svc *customerMemberService) Delete(id int64) error {
	return svc.repo.Delete(id)
}
