package customerGroupEntityRelation

type Repository interface {
	Get(id int64) (*CustomerGroupEntityRelation, error)
	GetAll() ([]*CustomerGroupEntityRelation, error)
	Save(custGER *CustomerGroupEntityRelation) error
	Delete(id int64) error
}
