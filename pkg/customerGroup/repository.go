package customerGroup

type Repository interface {
	// Get(id int64) (*CustomerGroup, error)
	Get(id int64, OffSet, FetchNext int) ([]*CustomerGroup, error)
	Get_Pagination(id int64) (int, error)
	Save(custGroup *CustomerGroup) error
	Delete(id int64) error
}
