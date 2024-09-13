package models

type CampaignModel struct {
	Id          string `gorm:"column:id;primaryKey" json:"id"`
	BranchId    string `gorm:"column:branch_id" json:"branchId"`
	StartDate   string `gorm:"column:start_date" json:"startDate"`
	EndDate     string `gorm:"column:end_date" json:"endDate"`
	BonusType   string `gorm:"column:bonus_type" json:"bonusType"`
	BonusValue  string `gorm:"column:bonus_value" json:"bonusValue"`
	MinPurchase string `gorm:"column:min_purchase" json:"minPurchase"`
	CreatedAt   string `gorm:"column:created_at" json:"createdAt"`
}
