package customerEntity

import "github.com/chandantechpanion/adminpanel-service/pkg/customerMaster"

//Service ...
type Service interface {
	// Get(id int64, Page, RecordsPerPage int) ([]*CustomerEntity, *pagination.Pagination, error)
	GetAll() ([]*CustomerEntity, error)
	Save(custEntity *CustomerEntity) error
	Get(id int64) ([]*CustomerEntity, error)
}

type customerEntityService struct {
	repo         Repository
	customerRepo customerMaster.Repository
}

//NewCustomerEntityService ...
func NewCustomerEntityService(repo Repository, customerRepo customerMaster.Repository) Service {
	return &customerEntityService{
		repo:         repo,
		customerRepo: customerRepo,
	}
}

func (svc *customerEntityService) GetAll() ([]*CustomerEntity, error) {
	return svc.repo.GetAll()
}

func (svc *customerEntityService) Save(custEntity *CustomerEntity) error {
	//Checking for Actually Customer master exists
	customerMaster, err := svc.customerRepo.GetByID(int(custEntity.CustomerMasterID))
	if err == nil && customerMaster.ID > 0 {
		return svc.repo.Save(custEntity)
	}
	return err
}

func (svc *customerEntityService) Get(id int64) ([]*CustomerEntity, error) {
	return svc.repo.Get(id)
}
