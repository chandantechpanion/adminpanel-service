package customerGroupEntityRelation

//Service ...
type Service interface {
	Get(id int64) (*CustomerGroupEntityRelation, error)
	GetAll() ([]*CustomerGroupEntityRelation, error)
	Save(custGER *CustomerGroupEntityRelation) error
	Delete(id int64) error
}

type customerGroupEntityRelationService struct {
	repo Repository
}

//NewCustomerGroupEntityRelationService ...
func NewCustomerGroupEntityRelationService(repo Repository) Service {
	return &customerGroupEntityRelationService{
		repo: repo,
	}
}

func (svc *customerGroupEntityRelationService) Get(id int64) (*CustomerGroupEntityRelation, error) {
	return svc.repo.Get(id)
}

func (svc *customerGroupEntityRelationService) GetAll() ([]*CustomerGroupEntityRelation, error) {
	return svc.repo.GetAll()
}

func (svc *customerGroupEntityRelationService) Save(custGER *CustomerGroupEntityRelation) error {
	return svc.repo.Save(custGER)
}

func (svc *customerGroupEntityRelationService) Delete(id int64) error {
	return svc.repo.Delete(id)
}
