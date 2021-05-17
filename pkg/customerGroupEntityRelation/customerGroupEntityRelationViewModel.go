package customerGroupEntityRelation

type CustomerGroupEntityRelationViewModel struct {
	CustomerMasterID int64  `json:"customer_master_id"`
	CustomerEntityID string `json:"customer_entity_id"`
	CustomerGroupID  int64  `json:"customer_group_id"`
}
