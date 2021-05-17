package customerMaster

type Repository interface {
	Get() ([]*CustomerMaster, error)
	GetByID(id int) (*CustomerMaster, error)
	Save(custMaster *CustomerMaster) error
	Delete(id int64) error
}
