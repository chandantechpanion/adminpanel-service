package customerGroup

import "github.com/chandantechpanion/adminpanel-service/pkg/pagination"

//Service ...
type Service interface {
	// Get(id int64) (*CustomerGroup, error)
	Get(id int64, Page, RecordsPerPage int) ([]*CustomerGroup, *pagination.Pagination, error)
	Save(custGroup *CustomerGroup) error
	Delete(id int64) error
}

type customerGroupService struct {
	repo Repository
}

//NewCustomerGroupService ...
func NewCustomerGroupService(repo Repository) Service {
	return &customerGroupService{
		repo: repo,
	}
}

// func (svc *customerGroupService) Get(id int64) (*CustomerGroup, error) {
// 	return svc.repo.Get(id)
// }

func (svc *customerGroupService) Get(id int64, Page, RecordsPerPage int) ([]*CustomerGroup, *pagination.Pagination, error) {

	pagination, errPag := svc.get_Pagination(id, Page, RecordsPerPage)

	if errPag != nil {
		return nil, nil, errPag
	} else {
		invoiceCaptureList, err := svc.repo.Get(id, pagination.Offset, pagination.Limit)
		if err != nil {
			return nil, nil, err
		} else {
			return invoiceCaptureList, pagination, nil
		}
	}
}

func (svc *customerGroupService) get_Pagination(userId int64, Page, RecordsPerPage int) (*pagination.Pagination, error) {

	TotalRows, err := svc.repo.Get_Pagination(userId)
	if err != nil {
		return nil, err
	}
	pagingData := pagination.GetPaginationData(Page, RecordsPerPage, TotalRows)

	return pagingData, nil
}

func (svc *customerGroupService) Save(custGroup *CustomerGroup) error {
	return svc.repo.Save(custGroup)
}

func (svc *customerGroupService) Delete(id int64) error {
	return svc.repo.Delete(id)
}
