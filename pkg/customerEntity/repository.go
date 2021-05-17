package customerEntity

type Repository interface {
	GetAll() ([]*CustomerEntity, error)
	Save(custEntity *CustomerEntity) error
	Get(id int64) ([]*CustomerEntity, error)
}
