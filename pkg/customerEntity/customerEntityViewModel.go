package customerEntity

type CustomerEntityViewModel struct {
	ID               string `json:"id"`
	CustomerMasterID int64  `json:"customer_master_id"`
	Name             string `json:"name"`
}
