package customerMember

type Repository interface {
	Get(id int64) (*CustomerMember, error)
	Save(custGroup *CustomerMember) error
	Delete(id int64) error
}
