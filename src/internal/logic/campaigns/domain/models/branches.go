package models

type BranchesModel struct {
	Id         string `json:"id"`
	MerchantId string `json:"merchant_id"`
	Name       string `json:"name"`
	Locations  string `json:"locations"`
	CreatedAt  string `json:"created_at"`
}
