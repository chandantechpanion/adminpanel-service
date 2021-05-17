package customerMaster

//Service ...
type Service interface {
	// Get(id int64, Page, RecordsPerPage int) ([]*CustomerMaster, *pagination.Pagination, error)
	Get() ([]*CustomerMaster, error)
	Save(custMaster *CustomerMaster) error
	Delete(id int64) error
}

type customerMasterService struct {
	repo Repository
}

//NewCustomerMasterService ...
func NewCustomerMasterService(repo Repository) Service {
	return &customerMasterService{
		repo: repo,
	}
}

func (svc *customerMasterService) Get() ([]*CustomerMaster, error) {
	return svc.repo.Get()
}

func (svc *customerMasterService) Save(custMaster *CustomerMaster) error {
	return svc.repo.Save(custMaster)
}

func (svc *customerMasterService) Delete(id int64) error {
	return svc.repo.Delete(id)
}
